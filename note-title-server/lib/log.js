var INFO = (function(){
    var timestamp = function(){};
    timestamp.toString = function(){
        return "[INFO " + (new Date).toLocaleString() + "]";    
    };

    return {
        log: console.log.bind(console, '%s', timestamp)
    }
})();

module.exports = INFO;