<script>
    import { apiUrl } from '../config';

    export let recordId;
    export let fileName;
    export let date_time;
    export let author;
</script>

<div class="flex items-center justify-center">
    <div class="w-full max-w-md border-2 border-black shadow-solid-primary bg-white rounded-lg p-2 m-4">
        <img class="rounded-md border-black border-2" src={ apiUrl + '/api/files/images/' + recordId + '/' + fileName } alt=""/>
        <p>{ author }</p>
        <p>{ date_time }</p>
    </div>
</div>