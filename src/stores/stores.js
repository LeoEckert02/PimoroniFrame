import { writable } from 'svelte/store';
import { apiUrl } from '../config';
import PocketBase from 'pocketbase';

const client = new PocketBase(apiUrl);

export const imageData = writable([]);
export const page = writable(1);

export const fetchImages = async (page) => {
    imageData.set([]);

    // fetch a paginated records list
    const resultList = await client.collection('images').getList(page, 5, {
        sort: '-created',
    });

    const loadedImages = resultList.items.map((data) => {
        return {
            id: data.id,
            image: data.image,
            created: data.created,
            author: data.author,
            date_time: data.created,
        }
    })

    imageData.set(loadedImages);
}