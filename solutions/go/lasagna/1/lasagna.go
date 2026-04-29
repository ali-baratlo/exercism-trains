package lasagna

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) (RemainTime int) {
    RemainTime= (OvenTime-actualMinutesInOven)
    return 
}

func PreparationTime(numberOfLayers int) (Preparation int){
    Preparation = (numberOfLayers * 2)
    return 
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
    
    Preparation := (numberOfLayers * 2)
    return actualMinutesInOven + Preparation
}
