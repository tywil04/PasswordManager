import * as base64 from "base64-arraybuffer"

import * as cryptography from "$lib/js/cryptography.js"


export const rootKey = "PasswordManager:"
export const authTokenKey = rootKey + "authToken"
export const databaseKeyKey = rootKey + "databaseKey"
export const vaultsKey = rootKey + "vaults"
export const passwordsKey = rootKey + "passwords"
export const notesKey = rootKey + "notes"


export async function getAuthToken() {
    return sessionStorage.getItem(authTokenKey) || undefined
}

export async function setAuthToken(value) {
    return sessionStorage.setItem(authTokenKey, value)
}

export async function getDatabaseKey() {
    const value = sessionStorage.getItem(databaseKeyKey) || undefined
    if (value !== undefined)
        return await cryptography.importDatabaseKey(base64.decode(value))
}

export async function setDatabaseKey(value) {
    const encodedValue = base64.encode(value)
    return sessionStorage.setItem(databaseKeyKey, encodedValue)
}


export async function getVaults() {
    const vaults = sessionStorage.getItem(vaultsKey) || undefined
    if (vaults !== undefined) 
        return JSON.parse(vaults)
}

export async function setVaults(vaults) {
    return sessionStorage.setItem(vaultsKey, JSON.stringify(vaults))
}

export async function addVaults(vaults) {
    return await setVaults(vaults.concat((await getVaults()) || []))
}

export async function removeVaults() {
    return sessionStorage.removeItem(vaultsKey)
}


export async function getPasswords() {
    const passwords = sessionStorage.getItem(passwordsKey) || undefined
    if (passwords !== undefined) 
        return JSON.parse(passwords)
}

export async function setPasswords(passwords) {
    return sessionStorage.setItem(passwordsKey, JSON.stringify(passwords))
}

export async function addPasswords(passwords) {
    Object.assign(passwords, passwords, await getPasswords())
    return await setPasswords(passwords || {})
}

export async function updatePassword(id, password) {
    let allPasswords = await getPasswords()
    let index = allPasswords.indexOf(allPasswords.filter((password) => password.id == id)[0])
    allPasswords[index] = password
    await setPasswords(allPasswords)
}

export async function removePassword(id) {
    let allPasswords = await getPasswords()
    await setPasswords(allPasswords.filter((password) => password.id != id))  
}

export async function removePasswords() {
    return sessionStorage.removeItem(passwordsKey)
}


export async function getNotes() {
    const notes = sessionStorage.getItem(notesKey) || undefined
    if (notes !== undefined) 
        return JSON.parse(notes)
}

export async function setNotes(notes) {
    return sessionStorage.setItem(notesKey, JSON.stringify(notes))
}

export async function addNotes(notes) {
    Object.assign(notes, notes, await getNotes())
    return await setPasswords(notes || {})
}

export async function updateNote(id, note) {
    let allNotes = await getNotes()
    let index = allNotes.indexOf(allNotes.filter((note) => note.id == id)[0])
    allNotes[index] = note
    await setNotes(allNotes)
}

export async function removeNote(id) {
    let allNotes = await getNotes()
    await setNotes(allNotes.filter((note) => note.id != id))  
}

export async function removeNotes() {
    return sessionStorage.removeItem(notesKey)
}

export async function clear() {
    return sessionStorage.clear()
}