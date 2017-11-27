/**
 * Created by zhangchf on 28/08/2017.
 */
package Hamming

fun compute(leftStrand: String, rightStand: String): Int {
    require(leftStrand.length == rightStand.length) {"leftStrand and rightStrand must be of equal length."}

    return leftStrand.zip(rightStand).count { (left, right) -> left != right }
}
