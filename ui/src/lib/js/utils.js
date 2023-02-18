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

export async function getJson(url, body, headers={}) {
    return await fetch(url, {
        method: "GET",
        headers: Object.assign({}, {
            "Content-type": "application/json",
        }, headers),
        body: JSON.stringify(body)
    })
}

export async function postJson(url, body={}, headers={}) {
    return await fetch(url, {
        method: "POST",
        headers: Object.assign({}, {
            "Content-type": "application/json",
        }, headers),
        body: JSON.stringify(body)
    })
}