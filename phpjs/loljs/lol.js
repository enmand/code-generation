//Do some setup for a Node environment
//Code fun begins here

var console = require("console");

if (!("function" === typeof printf)) {
    function printf(msg) {
        console.log(msg);
    };
}

class phpConsole {
    log(msg) {
        printf("%s\n", msg);
    }

};

if (!(undefined !== console) || console === 1) {
    console = new phpConsole();
}

function consolelog(msg) {
    console.log(msg);
};

var ar = ["this", "is", "an", "array", "of", "things", 1, true];

for (var elem of Object.values(ar)) {
    console.log(typeof elem);
}

class Something {
    constructor() {
        this.x = "hello";
        this.y = "world";
    }

    __constructor(ix = undefined, iy = undefined) {
        if (undefined !== ix) {
            this.x = ix;
        }

        if (undefined !== iy) {
            this.y = iy;
        }
    }

    getX() {
        return this.x;
    }

};

var s = new Something();
console.log(s.getX() + " " + s.y);