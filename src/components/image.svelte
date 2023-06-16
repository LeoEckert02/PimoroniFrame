<script>
    import {apiUrl} from '../config';
    import PocketBase from "pocketbase";
    import {onMount} from "svelte";

    const client = new PocketBase(apiUrl);

    onMount(async () => {
        let userRecord = await client.collection("users").getOne(author);
        userName = userRecord.name

        const dt = new Date(date_time);
        // Assuming the CET time zone is "Europe/Paris"
        const cet_timezone = 'Europe/Paris'; // Change this if needed
        const cet_dt = new Date(dt.toLocaleString('en-US', {timeZone: cet_timezone}));
        const options = {day: '2-digit', month: '2-digit', year: '2-digit', hour: '2-digit', minute: '2-digit'};
        cet_formatted_string = cet_dt.toLocaleString('en-GB', options).replace(',', '').replaceAll('/', '.').replace(' ', ' | ') + ' Uhr';

    })


    let userName;
    let cet_formatted_string
    export let recordId;
    export let fileName;
    export let date_time;
    export let author;
</script>

<div class="flex items-center justify-center">
    <div class="w-full max-w-md border-2 border-black shadow-solid-primary bg-white rounded-lg p-4 m-2 mx-4 mb-3">
        <img class="rounded-md border-black border-2 aspect-4/3"
             src={ apiUrl + '/api/files/images/' + recordId + '/' + fileName } alt=""/>
        <div class="rounded-md border-black border-2 p-2 mt-2 bg-white">
            <div class="ml-1">
                <p class="font-medium">{ userName }</p>
                <p class="text-gray-400 text-xs mb-1">{ cet_formatted_string }</p>
            </div>
        </div>
    </div>
</div>