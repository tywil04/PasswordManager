import * as base64 from "base64-arraybuffer"

import * as cryptography from "$lib/js/cryptography.js"

export const rootKey = "PasswordManager:"

export const authTokenKey = rootKey + "authToken"
export const databaseKeyKey = rootKey + "databaseKey"

export async function getAuthToken() {
    return sessionStorage.getItem(authTokenKey) || undefined
}

export function setAuthToken(value) {
    return sessionStorage.setItem(authTokenKey, value)
}

export async function getDatabaseKey() {
    const value = sessionStorage.getItem(databaseKeyKey) || undefined
    if (value !== undefined)
        return await cryptography.importDatabaseKey(base64.decode(value))
}

export function setDatabaseKey(value) {
    const encodedValue = base64.encode(value)
    return sessionStorage.setItem(databaseKeyKey, encodedValue)
}

export function clear() {
    sessionStorage.clear()
}