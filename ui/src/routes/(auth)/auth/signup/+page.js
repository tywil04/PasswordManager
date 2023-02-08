import { redirect } from "@sveltejs/kit";

import * as server from "$lib/js/server.js"

export async function load() {
    if (await server.isAuthed())
        throw redirect(302, "/")
}