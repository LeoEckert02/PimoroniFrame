<script>
    import Dropdown from '../components/dropdown.svelte'
    import Upload from '../components/upload.svelte'
    import {fetchImages} from "../stores/stores.js";
    import {onMount} from "svelte";
    import Image from "../components/image.svelte";
    import {imageData, page} from "../stores/stores";

    onMount(async () => {
        await fetchImages($page);
    });
</script>

<Dropdown/>
<Upload/>
{#each $imageData as image}
    <Image recordId={image.id} fileName={image.image} date_time={image.date_time} author={image.author}></Image>
{/each}
