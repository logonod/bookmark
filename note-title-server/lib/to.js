module.exports = function(promise) {
    return promise.then(data => {
            return [data, null];
        })
        .catch(err => [null, err]);
}