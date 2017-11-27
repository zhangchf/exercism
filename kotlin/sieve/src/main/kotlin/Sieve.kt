/**
 * Created by zhangchf on 30/08/2017.
 */

package Sieve

fun primesUpTo(limit: Int): List<Int> {
    println("""$$ 99.9""")

    var primeMap = HashMap<Int, Int>()
    (2..limit).fold(primeMap) {acc, i -> acc[i] = 0; acc }

    return (2..limit).fold(primeMap){
        acc, i ->
        (i..limit).onEach {
            j ->
            if (j != i && j % i == 0) acc[j] = 1
        }
        acc
    }.filterValues { it == 0 }.keys.toList()

}