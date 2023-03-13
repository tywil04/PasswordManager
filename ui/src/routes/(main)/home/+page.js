import { redirect } from "@sveltejs/kit";

import * as storage from "$lib/js/storage.js"
import * as server from "$lib/js/server.js"
import * as utils from "$lib/js/utils.js"
import * as cryptography from "$lib/js/cryptography.js"

import * as base64 from "base64-arraybuffer"

export async function load() {
    if (!await server.isAuthed())
        throw redirect(302, "/auth/signin")

    const vaults = await server.getVaults()
    const passwords = await server.getPasswords()
    const notes = await server.getNotes()

    return { vaults, passwords, notes }
}