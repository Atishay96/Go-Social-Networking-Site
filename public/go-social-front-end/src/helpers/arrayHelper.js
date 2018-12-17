
/**
 * 
 * @param {Object} object
 * @return {Array}
 */
export const notEmpty = (object) => {
    var empty = []

    Object.keys(object).map(v => {
        if (object[v] == '') {
            empty.push(v)
        }
    })
    return empty
}