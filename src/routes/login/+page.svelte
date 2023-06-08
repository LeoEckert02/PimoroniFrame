<script>
    import PocketBase from 'pocketbase';
    import toast from 'svelte-french-toast';
    import {goto} from "$app/navigation";
    import {onMount} from "svelte";

    const client = new PocketBase('http://127.0.0.1:8090');

    let username = '';
    let password = '';

    onMount(() => {
        if (client.authStore.isValid) {
            goto('/');
        }
    })

    function login() {
        toast.promise(auth(),
            {
                loading: 'Laden...',
                success: 'Login erfolgreich!',
                error: 'Login fehlgeschlagen!'
            },
            {
                position: 'bottom-center'
            },
        );

        setTimeout(() => {
            if (client.authStore.isValid) {
                goto('/');
            } else {
                username = '';
                password = '';
            }
        }, 300);
    }

    function auth() {
        return client.collection('users').authWithPassword(username, password);
    }

</script>

{#if !client.authStore.isValid }
    <div class="flex items-center justify-center h-screen">
        <div class="w-full max-w-sm m-4">
            <form on:submit|preventDefault={login}
                  class="bg-white border-2 border-black shadow-solid-primary rounded-lg px-8 pt-6 pb-8 mb-4 -mx3">
                <div class="text-3xl font-medium pb-3 text-black text-left">Login</div>
                <div class="mb-4">
                    <label class="block text-black text-md font-medium mb-2" for="username">
                        Benutzername
                    </label>
                    <input bind:value={username}
                           class="mb-1 border-2 border-black rounded-lg w-full py-3 px-4 shadow-solid-primary leading-tight"
                           id="username" type="text" required placeholder="Benutzername">
                </div>
                <div class="mb-6">
                    <label class="block text-black text-md font-medium mb-2" for="password">
                        Passwort
                    </label>
                    <input bind:value={password}
                           class="mb-1 appearance-none border-2 border-black rounded-lg w-full py-3 px-4 shadow-solid-primary leading-tight"
                           id="password" type="password" required placeholder="Passwort">
                </div>
                <button class="hover:bg-black hover:text-white font-medium bg-custom2 text-black items-center inline-flex bg-red-100 border-2 border-black duration-300 ease-in-out focus:outline-none hover:shadow-none justify-center rounded-lg shadow-solid-primary text-center transform transition w-full py-3 text-lg">
                    Login
                </button>
            </form>
        </div>
    </div>
{/if}
