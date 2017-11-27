import java.math.BigDecimal
import java.math.RoundingMode

/**
 * Created by zhangchf on 28/08/2017.
 */
class SpaceAge(val seconds: Long) {

    enum class Planet(val orbitalPeriod: Double) {
        EARTH(1.0),
        MERCURY(0.2408467),
        VENUS(0.61519726),
        MARS(1.8808158),
        JUPITER(11.862615),
        SATURN(29.447498),
        URANUS(84.016846),
        NEPTUNE(164.79132)
    }

    private val ORBITAL_PERIOD_EARTH = 31557600.0

    fun onEarth() = spaceAge(Planet.EARTH)
    fun onMercury() = spaceAge(Planet.MERCURY)
    fun onVenus() = spaceAge(Planet.VENUS)
    fun onMars() =  spaceAge(Planet.MARS)
    fun onJupiter() = spaceAge(Planet.JUPITER)
    fun onSaturn() = spaceAge(Planet.SATURN)
    fun onUranus() = spaceAge(Planet.URANUS)
    fun onNeptune() = spaceAge(Planet.NEPTUNE)


    private fun earthAgeInternal(): Double {
        return seconds / ORBITAL_PERIOD_EARTH
    }

    private fun spaceAge(planet: Planet): Double {
        return (earthAgeInternal() / planet.orbitalPeriod).roundTo(2)
    }

    companion object TimeUtils {
        fun Double.roundTo(decimalPlace: Int): Double = BigDecimal(this).setScale(decimalPlace, RoundingMode.HALF_UP).toDouble()
    }
}