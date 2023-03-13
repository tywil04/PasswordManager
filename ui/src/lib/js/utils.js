export function stringToArrayBuffer(string) {
    return new TextEncoder().encode(string)
}

export function arrayBufferToString(ab) {
    return new TextDecoder().decode(ab)
}
  
export function base64ToUint8(string) {
    return Uint8Array.from(atob(string), (char) => char.charCodeAt(0))
}

export function uint8ToBase64(uint8) {
    return btoa(new TextDecoder('utf8').decode(uint8))
}

export function arrayBufferToHex(byteArray) {
    return [...new Uint8Array(byteArray)].map(x => x.toString(16).padStart(2, '0')).join('');
}

export function hexToArrayBuffer(hex) {
    return (new Uint8Array(hex.match(/[\da-f]{2}/gi).map((h) => parseInt(h, 16)))).buffer
}

export function base64ToUrlBase64(text) {
    return text.replace(/\+/g, "\-").replace(/\//g, "_").replace(/\=/g, "")
}

export function base64UrlToBase64(text) {
    return text.replace(/\-/g, "\+").replace(/\_/g, "\/")
}

export async function getJson(url, data) {
    return fetchJson("GET", url, data)
}

export async function postJson(url, data) {
    return fetchJson("POST", url, data)
}

export async function deleteJson(url, data) {
    return fetchJson("DELETE", url, data)
}

export async function fetchJson(method, url, data) {
    data = data || {}
    
    if (data.headers === undefined) {
        data.headers = {}
    }

    data.headers["Content-Type"] = "application/json"

    const response = await fetch(url, {
        method: method, 
        headers: data.headers, 
        body: JSON.stringify(data.body), 
    })
    return { status: response.status, json: await response.json() }
}