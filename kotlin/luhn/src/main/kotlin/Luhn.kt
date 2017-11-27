/**
 * Created by zhangchf on 29/08/2017.
 */

object Luhn {

    fun isValid(input: String): Boolean {
        val serial = input.filter { it.isDigit() || it.isWhitespace() || return@isValid false }
        if (serial.trim().length < 2)  return false

        return serial
                .trim()
                .trimStart('0')
                .filter { it.isDigit() }
                .map { it.toInt() - '0'.toInt() }
                .mapIndexed { index, i -> if (index % 2 == 0) doubing(i) else i }
                .sum()
                .rem(10) == 0
    }

    fun doubing(t: Int): Int {
        var result = t * 2
        return if (result > 9) result - 9 else result
    }

}