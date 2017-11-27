/**
 * Created by zhangchf on 29/08/2017.
 */

object SumOfMultiples {

    fun sum(factorSet: Set<Int>, max: Int): Int =
        (1 until max).filter { num -> factorSet.any { num.rem(it) == 0 } }.sum()

}
