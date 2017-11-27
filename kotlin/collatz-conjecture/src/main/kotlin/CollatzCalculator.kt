/**
 * Created by zhangchf on 31/08/2017.
 */

object CollatzCalculator {
/*    fun computeStepCount(number: Int): Int {
        require(number > 0) {"Only natural numbers are allowed"}

        tailrec fun compute(number: Int, step: Int): Int =
                when {
                    number == 1 -> step
                    number % 2 == 1 -> compute(3 * number + 1, step + 1)
                    else -> compute(number / 2, step + 1)
                }

        return compute(number, 0)
    }*/

    fun computeStepCount(number: Int): Int {

        require(number > 0) { "Only natural numbers are allowed" }
        return generateSequence(number) { if (it % 2 == 0) it / 2 else it * 3 + 1 }
                .takeWhile { it != 1 }
                .count()
    }
}