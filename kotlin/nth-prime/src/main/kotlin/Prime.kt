/**
 * Created by zhangchf on 04/09/2017.
 */

object Prime {
    var primes = mutableListOf(2, 3)

    fun nth(index: Int): Int {
        require(index > 0)

        tailrec fun findPrime(index: Int): Int {
            return when {
                primes.size >= index -> primes.get(index - 1)
                else -> {
                    val next = generateSequence(primes.last()) {it + 1}.find { num -> primes.all { num % it != 0 } }
                    primes.add(next!!)
                    findPrime(index)
                }
            }
        }
        return findPrime(index)
    }
}