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
                loading: 'Loading...',
                success: 'Login successful!',
                error: 'Invalid credentials!'
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
        return client.admins.authWithPassword(username, password);
    }

</script>

{#if !client.authStore.isValid }
    <div class="flex items-center justify-center h-screen">
        <div class="w-full max-w-xs">
            <form on:submit|preventDefault={login} class="bg-white shadow-lg rounded-md px-8 pt-6 pb-8 mb-4">
                <div class="text-2xl font-bold pb-3 text-gray-600 text-left">Admin Login</div>
                <div class="mb-4">
                    <label class="block text-gray-500 text-sm font-bold mb-2" for="username">
                        Username
                    </label>
                    <input bind:value={username} class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500" id="username" type="text" placeholder="Username">
                </div>
                <div class="mb-6">
                    <label class="block text-gray-500 text-sm font-bold mb-2" for="password">
                        Password
                    </label>
                    <input bind:value={password} class="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-blue-500" id="password" type="password" placeholder="Password">
                </div>
                <div class="flex items-center justify-between">
                    <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
                        Login
                    </button>
                </div>
            </form>
        </div>
    </div>
{/if}
