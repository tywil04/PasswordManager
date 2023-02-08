import * as base64 from "base64-arraybuffer"

export const rootKey = "PasswordManager:"

export const authTokenKey = rootKey + "authToken"
export const databaseKeyKey = rootKey + "databaseKey"

export function getAuthToken() {
    return sessionStorage.getItem(authTokenKey) || undefined
}

export function setAuthToken(value) {
    return sessionStorage.setItem(authTokenKey, value)
}

export function getDatabaseKey() {
    const value = sessionStorage.getItem(databaseKeyKey) || undefined
    if (value !== undefined)
        return base64.decode(value)
}

export function setDatabaseKey(value) {
    const encodedValue = base64.encode(value)
    return sessionStorage.setItem(databaseKeyKey, encodedValue)
}

export function clear() {
    sessionStorage.clear()
}