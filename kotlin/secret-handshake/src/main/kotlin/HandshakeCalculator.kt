/**
 * Created by zhangchf on 29/08/2017.
 */

object HandshakeCalculator {

    fun calculateHandshake(input:Int): List<Signal> {

        var signals = ArrayList<Signal>()

        (0..4).filter { input and (1 shl it) != 0 }.map { if (it < 4) signals.add(Signal.values()[it]) else signals.reverse() }

        return signals
    }
}