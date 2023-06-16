<script>
    let hidden = true;

    function onClick() {
        hidden = !hidden;
    }
</script>

<div class="md:fixed md:z-90 md:top-3 md:right-1/4 text-right p-3 -mb-2">
    <div class="md:relative inline-block text-left">
        <div>
            <button on:click={onClick} class="hover:bg-black hover:text-white bg-custom2 w-10 h-10 rounded-full shadow-solid-secondary border-2 border-black flex justify-center items-center duration-300 ease-in-out focus:outline-none hover:shadow-none">
                <svg fill="currentColor" xmlns="http://www.w3.org/2000/svg" height="1.2em" viewBox="0 0 448 512">
                    <!--! Font Awesome Free 6.4.0 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license (Commercial License) Copyright 2023 Fonticons, Inc. -->
                    <path d="M0 96C0 78.3 14.3 64 32 64H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H32C14.3 128 0 113.7 0 96zM0 256c0-17.7 14.3-32 32-32H416c17.7 0 32 14.3 32 32s-14.3 32-32 32H32c-17.7 0-32-14.3-32-32zM448 416c0 17.7-14.3 32-32 32H32c-17.7 0-32-14.3-32-32s14.3-32 32-32H416c17.7 0 32 14.3 32 32z"/>
                </svg>
            </button>
        </div>

        {#if !hidden}
            <div on:click={onClick} class="m-3 mb-0 font-medium absolute right-0 z-10 mt-2 w-56 origin-top-right divide-y-2 divide-black rounded-lg bg-white shadow-solid-primary border-black border-2"
                 role="menu" aria-orientation="vertical" aria-labelledby="menu-button" tabindex="-1">
                <div class="hover:bg-gray-100 py-1 rounded-t-lg" role="none">
                    <!-- Active: "bg-gray-100 text-gray-900", Not Active: "text-gray-700" -->
                    <a href="/settings" class="text-gray-700 block px-4 py-2 text-sm" role="menuitem" tabindex="-1"
                       id="menu-item-0">
                        Einstellungen
                    </a>
                </div>
                <div class="hover:bg-gray-100 py-1" role="none">
                    <a href="/datenschutz" class="text-gray-700 block px-4 py-2 text-sm" role="menuitem" tabindex="-1"
                       id="menu-item-2">Datenschutzerklärung</a>
                </div>
                <div class="hover:bg-gray-100 py-1 rounded-b-lg" role="none">
                    <a href="/logout" class="text-gray-700 block px-4 py-2 text-sm" role="menuitem" tabindex="-1"
                       id="menu-item-4">Logout</a>
                </div>
            </div>
        {/if}
    </div>
</div>
