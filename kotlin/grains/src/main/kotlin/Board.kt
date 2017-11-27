import java.math.BigInteger

/**
 * Created by zhangchf on 30/08/2017.
 */

object Board {

    var grainMap = hashMapOf<Int, BigInteger>(1 to BigInteger.valueOf(1L))

    fun getGrainCountForSquare(index: Int): BigInteger {
        require(index > 0 && index <= 64) {"Only integers between 1 and 64 (inclusive) are allowed"}
        if (grainMap.get(index) == null) {
            grainMap[index] = getGrainCountForSquare(index - 1).shiftLeft(1).
        }
        return grainMap[index]!!
    }

    fun getTotalGrainCount(): BigInteger {
        return grainMap.values.fold(BigInteger.valueOf(0L)) { acc, l -> acc + l }
    }

}