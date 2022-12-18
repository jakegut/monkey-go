let map = fn(arr, f){
    let iter = fn(arr, acc){
        if (len(arr) == 0){
            acc
        } else {
            iter(rest(arr), push(acc, f(first(arr))));
        }
    }

    iter(arr, [])
};

let range = fn(start, end) {
    let iter = fn(i, acc){
        if (i == end + 1) {
            return acc
        } else {
            iter(i + 1, push(acc, i))
        }
    }

    iter(start, [])
}

let a = [1, 2, 3, 4];
let double = fn(x) { x * 2 };
let doubled = map(a, double)


