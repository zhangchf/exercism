/**
 * Created by zhangchf on 29/08/2017.
 */
class Squares(val number: Int) {
    fun squareOfSum() = sequence().sum().squared()

    fun sumOfSquares() = sequence().map { it.squared() }.sum()

    fun difference() = squareOfSum() - sumOfSquares()

    fun sequence() = 1 .. number

    fun Int.squared(): Int = this * this
}