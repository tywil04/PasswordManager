import * as storage from "$lib/js/storage.js"
import * as utils from "$lib/js/utils.js"
import * as cryptography from "$lib/js/cryptography.js"

import * as base64 from "base64-arraybuffer"


export async function isAuthed() {
    const authToken = await storage.getAuthToken()

    if (authToken == undefined) {
        return false
    }

    const { status } = await utils.getJson("/api/v1/auth/test", {
        headers: { 
            "Authorization": authToken
        }
    })

    if (status === 200) {
        return true
    }

    storage.clear()
    return false
}

export async function signout() {
    let { status } = await utils.deleteJson("/api/v1/auth/signout", { 
        headers: {
            "Authorization": await storage.getAuthToken()
        }
    })

    if (status === 200) {
        sessionStorage.clear()
        location.reload()
    }
}

export async function syncClientData() {
    const tempVaults = await storage.getVaults()
    const tempPasswords = await storage.getPasswords()
    const tempNotes = await storage.getNotes()

    await storage.removeVaults()
    await storage.removePasswords()
    await storage.removeNotes()

    if (await getVaults() === undefined)
        storage.setVaults(tempVaults)

    if (await getPasswords() === undefined)
        storage.setPasswords(tempPasswords)

    if (await getNotes() === undefined)
        storage.setNotes(tempNotes)

    window.location.reload()
}

export async function getVaults() {
    const authToken =  await storage.getAuthToken()
    const databaseKey = await storage.getDatabaseKey()

    let vaults = await storage.getVaults()
    if (vaults === undefined) {
        const response = (await utils.getJson("/api/v1/vaults", { 
            headers: { 
                "Authorization": authToken
            } 
        }))
        
        if (response.status !== 200)
            return undefined
        
        vaults = response.json.vaults

        for (let vault of vaults) {
            vault.name = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(vault.name),
                iv: base64.decode(vault.nameIv),
            })) 
    
            vault.colour = "#" + utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(vault.colour),
                iv: base64.decode(vault.colourIv),
            }))     
    
            delete vault.nameIv
            delete vault.colourIv
        }

        await storage.setVaults(vaults)
    }

    return vaults
}

export async function getPasswords() {
    const authToken =  await storage.getAuthToken()
    const databaseKey = await storage.getDatabaseKey()

    const parsedPasswords = {}
    let passwords = await storage.getPasswords()

    if (passwords === undefined) {
        const response = (await utils.getJson(`/api/v1/vaults/passwords`, { 
            headers: { 
                "Authorization": authToken,
            }
        }))

        if (response.status !== 200)
            return undefined
        
        passwords = response.json.passwords

        for (let password of passwords) {
            password.name = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(password.name),
                iv: base64.decode(password.nameIv),
            })) 
    
            password.colour = "#" + utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(password.colour),
                iv: base64.decode(password.colourIv),
            }))     
            
            delete password.nameIv 
            delete password.colourIv
        }

        await storage.setPasswords(passwords)
    }
        
    for (let password of passwords) {
        password.password = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
            encrypted: base64.decode(password.password),
            iv: base64.decode(password.passwordIv),
        })) 

        password.username = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
            encrypted: base64.decode(password.username),
            iv: base64.decode(password.usernameIv),
        }))

        delete password.passwordIv
        delete password.usernameIv

        for (let additionalField of password.additionalFields) {
            additionalField.key = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(additionalField.key),
                iv: base64.decode(additionalField.keyIv),
            })) 

            additionalField.value = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(additionalField.value),
                iv: base64.decode(additionalField.valueIv),
            })) 
            
            delete additionalField.keyIv
            delete additionalField.valueIv
        }

        for (let url of password.urls) {
            url.url = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(url.url),
                iv: base64.decode(url.urlIv),
            })) 

            delete url.urlIv
        }

        const passwordId = password.id 
        delete password.id 

        parsedPasswords[passwordId] = password
    }

    return parsedPasswords
}

export async function getNotes() {
    const authToken =  await storage.getAuthToken()
    const databaseKey = await storage.getDatabaseKey()

    const parsedNotes = {}
    let notes = await storage.getNotes()

    if (notes === undefined) {
        const response = (await utils.getJson(`/api/v1/vaults/notes`, { 
            headers: { 
                "Authorization": authToken,
            }
        }))

        if (response.status !== 200)
            return undefined
        
        notes = response.json.notes

        for (let note of notes) {
            note.name = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(note.name),
                iv: base64.decode(note.nameIv),
            })) 
    
            note.colour = "#" + utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                encrypted: base64.decode(note.colour),
                iv: base64.decode(note.colourIv),
            }))       
            
            delete note.nameIv 
            delete note.colourIv
        }

        await storage.setNotes(notes)
    }

    for (let note of notes) {
        note.title = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
            encrypted: base64.decode(note.title),
            iv: base64.decode(note.titleIv),
        })) 

        note.content = utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
            encrypted: base64.decode(note.content),
            iv: base64.decode(note.contentIv),
        }))

        delete note.titleIv
        delete note.contentIv

        const noteId = note.id 
        delete note.id 

        parsedNotes[noteId] = note
    }

    return parsedNotes
}

export async function updatePassword(id, newPassword) {
    const databaseKey = await storage.getDatabaseKey()
    const authToken = await storage.getAuthToken()

    const encryptedName = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(newPassword.name))
    const encryptedUsername = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(newPassword.username))
    const encryptedPassword = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(newPassword.password))
    const encryptedColour = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(newPassword.colour.replace("#", "")))

    const processedAdditionalFields = []
    const processedUrls = []

    for (let additionalField of newPassword.additionalFields) {
        const encryptedKey = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(additionalField.key))
        const encryptedValue = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(additionalField.value))

        processedAdditionalFields.push({
            key: base64.encode(encryptedKey.encrypted),
            keyIv: base64.encode(encryptedKey.iv),
            value: base64.encode(encryptedValue.encrypted),
            valueIv: base64.encode(encryptedValue.iv)
        })
    }

    for (let url of newPassword.urls) {
        const encryptedUrl = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(url.url))

        processedUrls.push({
            url: base64.encode(encryptedUrl.encrypted),
            urlIv: base64.encode(encryptedUrl.iv)
        })
    }

    let updatedPassword = {
        passwordId: id,
        vaultId: newPassword.vaultId,
        name: base64.encode(encryptedName.encrypted),
        nameIv: base64.encode(encryptedName.iv),
        username: base64.encode(encryptedUsername.encrypted),
        usernameIv: base64.encode(encryptedUsername.iv),
        password: base64.encode(encryptedPassword.encrypted),
        passwordIv: base64.encode(encryptedPassword.iv),
        colour: base64.encode(encryptedColour.encrypted),
        colourIv: base64.encode(encryptedColour.iv),
        additionalFields: processedAdditionalFields,
        urls: processedUrls,
    }

    const { status, json } = await utils.putJson("/api/v1/vaults/passwords", {  
        body: updatedPassword,
        headers: {
            "Authorization": authToken,
        }
    })

    if (status == 200) {
        updatedPassword.name = newPassword.name
        updatedPassword.colour = newPassword.colour
    
        delete updatedPassword.nameIv
        delete updatedPassword.colourIv

        await storage.updatePassword(id, updatedPassword)
    }

    return { status, json }
}

export async function updateNote(id, newNote) {
    const databaseKey = await storage.getDatabaseKey()
    const authToken = await storage.getAuthToken()

    const encryptedName = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(newNote.name))
    const encryptedTitle = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(newNote.title))
    const encryptedContent = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(newNote.content))
    const encryptedColour = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(newNote.colour.replace("#", "")))

    let updatedNote = {
        noteId: id,
        vaultId: newNote.vaultId,
        name: base64.encode(encryptedName.encrypted),
        nameIv: base64.encode(encryptedName.iv),
        title: base64.encode(encryptedTitle.encrypted),
        titleIv: base64.encode(encryptedTitle.iv),
        content: base64.encode(encryptedContent.encrypted),
        contentIv: base64.encode(encryptedContent.iv),
        colour: base64.encode(encryptedColour.encrypted),
        colourIv: base64.encode(encryptedColour.iv),
    }

    const { status, json } = await utils.putJson("/api/v1/vaults/notes", {  
        body: updatedNote,
        headers: {
            "Authorization": authToken,
        }
    })

    if (status == 200) {
        updatedNote.name = newPassword.name
        updatedNote.colour = newPassword.colour
    
        delete updatedNote.nameIv
        delete updatedNote.colourIv

        await storage.updateNote(id, updatedNote)
    }

    return { status, json }
}

export async function deletePassword(id, vaultId) {
    const authToken = await storage.getAuthToken()

    const { status, _ } = await utils.deleteJson("/api/v1/vaults/passwords", {  
        body: {
            passwordId: id,
            vaultId: vaultId,
        },
        headers: {
            "Authorization": authToken,
        }
    })

    if (status == 200) {
        await storage.removePassword(id)
    }
}

export async function deleteNote(id, vaultId) {
    const authToken = await storage.getAuthToken()

    const { status, _ } = await utils.deleteJson("/api/v1/vaults/notes", {  
        body: {
            noteId: id,
            vaultId: vaultId,
        },
        headers: {
            "Authorization": authToken,
        }
    })

    if (status == 200) {
        await storage.removeNote(id)
    }
}