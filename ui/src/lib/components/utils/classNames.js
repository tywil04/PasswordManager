export function classNames(...classNames) {
    let result = ""
    for (let className of classNames) {
        result += className + " "
    }
    return result.substring(0, result.length - 1)
}