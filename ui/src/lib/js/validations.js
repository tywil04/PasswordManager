export function validateEmail(email) {
    const emailRegex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    return emailRegex.test(email)
}

export function validatePassword(password) {
    const passwordRegex = /^((?=.*([A-Z]){1,})(?=.*([!#$%^'"`&*-=_+><?;:(){}\[\].,@]){1,})(?=.*([0-9]){1,})(?=.*([a-z]){1,})).{8,}$/
    return passwordRegex.test(password)
}

export function validateHexColour(colour) {
    const colourRegex = /#[0123456789abcdef]{6,6}$/
    return colourRegex.test(colour)
}