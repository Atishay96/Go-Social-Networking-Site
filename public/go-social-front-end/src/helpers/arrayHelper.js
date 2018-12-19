
/**
 * 
 * @param {Object} object
 * @return {Array}
 */
export const notEmpty = (object) => {
    var empty = []
    Object.keys(object).map(v => {
        if (object[v] === '' || object[v] === false || object[v] === null || object[v] === undefined) {
            empty.push(v)
        }
        return v
    })
    return empty
}