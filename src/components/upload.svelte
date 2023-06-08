<script>
    import {goto} from "$app/navigation";
    import PocketBase from "pocketbase";
    import {onMount} from "svelte";
    import {fade} from 'svelte/transition';
    import toast from "svelte-french-toast";

    const client = new PocketBase('http://127.0.0.1:8090');

    onMount(() => {
        if (!client.authStore.isValid) {
            goto('/login');
        }
    })

    let input;
    let image;
    let imageFile;
    let showImage = false;

    function onChange() {
        const file = input.files[0];

        console.log((file.name.split('.').pop()));

        if (file) {
            if (file.name.split('.').pop().toLowerCase() === 'heic') {
                // Create a new FormData object
                const formData = new FormData();

                // Append the selected file to the FormData object
                formData.append('image', file);

                // Send the HTTP request with the FormData
                fetch('http://127.0.0.1:8090/api/convert', {
                    method: 'PUT',
                    body: formData
                }).then(response => {
                    // Check the response status
                    if (response.ok) {
                        // Read the response data as a blob
                        return response.blob();
                    } else {
                        throw new Error('Request failed with status ' + response.status);
                    }
                })
                    .then(blob => {
                        // Create a URL object from the blob
                        const imageUrl = URL.createObjectURL(blob);

                        // Use the imageUrl as the source for an <img> element or perform other operations
                        image.setAttribute("src", imageUrl)
                        imageFile = new File([blob], file.name.split('.')[0] + '.jpg'); // Create a File object from the blob
                        console.log(imageFile.size);
                    })
                    .catch(error => {
                        // Handle any errors that occurred during the request
                        console.error('Error:', error);
                    });
            } else {
                const reader = new FileReader();
                reader.addEventListener("load", function () {
                    image.setAttribute("src", reader.result);
                });
                reader.readAsDataURL(file);
                imageFile = file;
            }

            showImage = true;

            return;
        }
        showImage = false;
    }

    function onUpload() {
        toast.promise(upload(),
            {
                loading: 'Hochladen...',
                success: 'Hochgeladen!',
                error: 'Hochladen fehlgeschlagen!'
            },
            {
                position: 'bottom-center'
            },
        );
    }

    function upload() {
        if (image) {
            let formData = new FormData();
            formData.append("author", client.authStore.model.id);
            formData.append("image", imageFile);

            return client.collection('images').create(formData);
        }

        toast.error('Wähle ein Bild aus!', {
            position: "bottom-center"
        })
    }

    function otherImg() {
        showImage = false;
        input.value = '';
        console.log(input);
        console.log(image);
        console.log(imageFile);
    }
</script>

{#if client.authStore.isValid}
    <div class="flex items-center justify-center h-screen">
        <div class="w-full max-w-md border-2 border-black shadow-solid-primary bg-white rounded-lg p-2 m-4">
            <div class="px-6 pt-4 lg:pb-3 pb-1">
                <div class="font-medium text-2xl mb-2 text-custom-1">Teile deine Erinnerungen mit Großmutter!</div>
                <h3 class="text-lg pb-1 font-medium">Wie funktionierts?</h3>
                <p class="text-gray-700 text-base">
                    Drücke einfach auf die Bild-Auswahl-Fläche und lade dein gewünschtes Bild hoch.
                    Sobald es hochgeladen ist, wird es mit etwas Verzögerung auf dem digitalen Bilderrahmen bei
                    Großmutter zu Hause angezeigt.
                </p>
            </div>
            <label for="dropzone-file"
                   class="hover:cursor-pointer md:p-5 p-3 m-4 mx-6 mb-4 aspect-4/3 flex items-center justify-center border-2 border-black rounded-lg shadow-solid-primary cursor-pointer bg-dotted-spacing-2 bg-dotted-gray-300">
                <div class="flex flex-col items-center justify-center">
                    {#if showImage}
                        <img in:fade class="rounded-md border-black border-2" bind:this={image} src="" alt=""/>
                    {:else}
                        <div class="border-2 border-black mb-3 bg-custom2 rounded-full">
                            <svg aria-hidden="true" class="w-11 h-11 m-3" fill="none" stroke="currentColor"
                                 viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                      d="M7 16a4 4 0 01-.88-7.903A5 5 0 1115.9 6L16 6a5 5 0 011 9.9M15 13l-3-3m0 0l-3 3m3-3v12"></path>
                            </svg>
                        </div>
                        <p class="text-center mb-2 text-sm"><span class="font-semibold">Klicke zum Hochladen</span> oder
                            ziehe das Bild hierher</p>
                    {/if}
                </div>
                <input bind:this={input} on:change={onChange} id="dropzone-file" type="file" class="hidden"/>
            </label>
            {#if showImage}
                <div class="m-6 mt-2 grid grid-cols-2">
                    <button on:click={otherImg} class="hover:bg-black hover:text-white mr-2 font-medium text-black items-center inline-flex border-2 border-black duration-300 ease-in-out focus:outline-none hover:shadow-none justify-center rounded-lg shadow-solid-primary text-center transform transition py-2 md:py-3">
                        <div class="flex items-center justify-center">
                            <span class="inline-flex items-center justify-center font-medium">
                                  Anderes Bild
                            </span>
                        </div>
                    </button>
                    <button on:click={onUpload} class="hover:bg-black hover:text-white font-medium bg-custom2 text-black items-center inline-flex border-2 border-black duration-300 ease-in-out focus:outline-none hover:shadow-none justify-center rounded-lg shadow-solid-primary text-center transform transition py-3">
                        <div class="flex items-center justify-center">
                            <span class="inline-flex items-center font- mr-1">
                              Hochladen
                            </span>
                            <svg fill="currentColor" width="18" height="18" xmlns="http://www.w3.org/2000/svg"
                                 viewBox="0 0 384 512">
                                <!--! Font Awesome Pro 6.4.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
                                <path d="M64 0C28.7 0 0 28.7 0 64V448c0 35.3 28.7 64 64 64H320c35.3 0 64-28.7 64-64V160H256c-17.7 0-32-14.3-32-32V0H64zM256 0V128H384L256 0zM216 408c0 13.3-10.7 24-24 24s-24-10.7-24-24V305.9l-31 31c-9.4 9.4-24.6 9.4-33.9 0s-9.4-24.6 0-33.9l72-72c9.4-9.4 24.6-9.4 33.9 0l72 72c9.4 9.4 9.4 24.6 0 33.9s-24.6 9.4-33.9 0l-31-31V408z"/>
                            </svg>
                        </div>
                    </button>
                </div>
            {/if}
        </div>
    </div>
{/if}