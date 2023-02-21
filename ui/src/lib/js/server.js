import * as storage from "$lib/js/storage.js"

export async function isAuthed() {
    const status = (await fetch("/api/v1/auth/test", {
        method: "GET",
        headers: {
            "Content-type": "application/json",
            "Authorization": await storage.getAuthToken()
        }
    })).status

    if (status === 200) {
        return true
    } else {
        storage.clear()
        return false
    }
}

export async function signout() {
    await fetch("/api/v1/auth/signout", {
        method: "DELETE",
        headers: {
            "Content-type": "application/json",
            "Authorization": sessionStorage.getItem("PasswordManager:authToken"),
        }
    })
    sessionStorage.clear()
    location.reload()
}