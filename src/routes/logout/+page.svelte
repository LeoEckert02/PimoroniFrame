<script>
    import {onMount} from "svelte";
    import PocketBase from "pocketbase";
    import {goto} from "$app/navigation";

    const client = new PocketBase('http://127.0.0.1:8090');

    onMount(() => {
        client.authStore.clear();
        goto('/login');
    })
</script>