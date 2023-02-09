import { redirect } from "@sveltejs/kit";

import * as storage from "$lib/js/storage.js"
import * as server from "$lib/js/server.js"

export async function load() {
    if (!await server.isAuthed())
        throw redirect(302, "/auth/signin")

    const json = await (await fetch("/api/v1/password", {
        method: "GET",
        headers: {
            "Content-type": "application/json",
            "Authorization": await storage.getAuthToken(),
        }
    })).json()

    return { passwords: json.passwords }
}