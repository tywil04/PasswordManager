export function validateEmail(email) {
    let emailRegex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
    return emailRegex.test(email)
}

export function validatePassword(password) {
    let passwordRegex = /^((?=.*([A-Z]){2,})(?=.*([!#$%^'"`&*-=_+><?;:(){}\[\].,@]){2,})(?=.*([0-9]){2,})(?=.*([a-z]){2,})).{12,}$/
    return passwordRegex.test(password)
}