package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	name   string
	power  int
	health int
	money  int
}

func main() {
	Hero1 := NewHero(reg())
	Greetings(Hero1.name)
	fightWithMobNew(Hero1.name, Hero1.power, Hero1.health, Hero1.money)

}

func Greetings(name string) {
	fmt.Printf("Добро пожаловать в симулятор говна, %s.\n Ты готов начать?\n Введи ДА или НЕТ: ", name)
	var answer string
	fmt.Scan(&answer)
	switch {
	case answer == "ДА":
		fmt.Println("Готовься к приключениям!")
		time.Sleep(1 * time.Second)

	case answer == "НЕТ":
		fmt.Println("Ну тебя тут и не ждали, мразь")
		panic("Мразь")
	default:
		fmt.Println("Ответь как положено, героическая жопа")
		Greetings(name)
	}
}

func randGeneratorPower() int {
	rand.Seed(time.Now().UnixNano())
	m := randInt(10, 30)
	return m
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func NewHero(name string, power, health int) Hero {
	return Hero{
		name:   name,
		power:  power,
		health: health,
	}
}

func reg() (string, int, int) {
	var name string
	fmt.Println("Введите имя героя:")
	fmt.Scan(&name)
	return name, randGeneratorPower(), 100
}

func orcNameGenerator() string {
	orcNames := []string{"Говномес", "Бульварный гей", "Ваня Захаров", "Супер зелёный гей гленомес", "Членозубый мудила"}
	return orcNames[randInt(1, len(orcNames))]
}

func orcGenerator(heroHealth, heroPower int) Hero {
	Orc := NewHero(orcNameGenerator(), randInt(1, 10)+heroHealth/2, 30+heroPower/4)
	time.Sleep(5 * time.Second)
	return Orc
}
func fightWithMobNew(name string, power, health, money int) {
	var (
		answer string
	)
	s := fmt.Sprintf("\nПривет, %s!\nВаши характеристики:\nСила: %d\nЗдоровье:%d.\nУ тебя в кармане %d монет.", name, power, health, money)
	fmt.Println(s)
	Orc := orcGenerator(health, power)
	time.Sleep(5 * time.Second)
	m := fmt.Sprintf("Вам повстречался орк по имени %s!\n Эта ленивая жопа обладает силой %d и здоровьем %d! Будем сражаться или поищем другого урода?\n Варианты ответа:\nБой\nМагазин\nХочу другого орка", Orc.name, Orc.power, Orc.health)
	fmt.Println(m)
	fmt.Scan(&answer)
	if answer == "Бой" {
		time.Sleep(5 * time.Second)
		fightWithOrc(Orc.name, name, Orc.power, power, Orc.health, health, money)
	} else if answer == "Магазин" {
		goToShop(name, health, money, power)
	} else if answer == "Хочу другого орка" {
		fightWithMobNew(name, power, health, money)
	}
}

func fightWithOrc(orcName, heroName string, oPower, hPower, oHealth, hHealth, hMoney int) {
	var (
		init     int
		choise   string
		punch    int
		orcPunch int
	)
	hHealthStatic := hHealth
	init = randInt(0, 10)
	if oHealth > 0 && hHealth > 0 {
		if init >= 6 {
			halfPowerHero := hPower/2 - randInt(0, 10)
			fmt.Println("Удача на твоей стороне! Ты успеваешь ударить быстрее вонючего орка!")
			fmt.Println("Чтобы совершить успешную атаку, направь свой удар!")
			fmt.Println("Напиши, в какую сторону ты его направишь!")
			fmt.Println("Варианты: Влево, Вправо")
			fmt.Scan(&choise)
			punch = randInt(0, 1)
			if choise == "Вправо" && punch == 1 {
				time.Sleep(3 * time.Second)
				oHealth = oHealth - hPower
				x1 := fmt.Sprintf("Отличная атака!\n Ты ударил на полную силу и нанёс %d урона!\n У него осталось %d здоровья", hPower, oHealth)
				fmt.Println(x1)
				fightWithOrc(orcName, heroName, oPower, hPower, oHealth, hHealth, hMoney)
			} else if choise == "Вправо" && punch == 0 {
				if halfPowerHero > 0 {
					time.Sleep(3 * time.Second)
					oHealth = oHealth - halfPowerHero
					x2 := fmt.Sprintf("Хорошая попытка, но орк ловко увернулся, однако ты смог нанести ему немного урона в размере %d\n У него осталось %d здоровья", halfPowerHero, oHealth)
					fmt.Println(x2)
					fightWithOrc(orcName, heroName, oPower, hPower, oHealth, hHealth, hMoney)
				} else {
					time.Sleep(3 * time.Second)
					fmt.Println("Неудача")
					fightWithOrc(orcName, heroName, oPower, hPower, oHealth, hHealth, hMoney)
				}
			}
			if choise == "Влево" && punch == 0 {
				time.Sleep(3 * time.Second)
				oHealth = oHealth - hPower
				x1 := fmt.Sprintf("Отличная атака!\n Ты ударил на полную силу и нанёс %d урона!\n У него осталось %d здоровья", hPower, oHealth)
				fmt.Println(x1)
				time.Sleep(3 * time.Second)
				fightWithOrc(orcName, heroName, oPower, hPower, oHealth, hHealth, hMoney)
			} else if choise == "Вправо" && punch == 1 {
				if halfPowerHero > 0 {
					oHealth = oHealth - halfPowerHero
					x2 := fmt.Sprintf("Хорошая попытка, но орк ловко увернулся, однако ты смог нанести ему немного урона в размере %d\n У него осталось %d здоровья.", halfPowerHero, oHealth)
					fmt.Println(x2)
					time.Sleep(3 * time.Second)
					fightWithOrc(orcName, heroName, oPower, hPower, oHealth, hHealth, hMoney)
				}
				fmt.Println("Неудача!")
				time.Sleep(3 * time.Second)
				fightWithOrc(orcName, heroName, oPower, hPower, oHealth, hHealth, hMoney)
			}

		}
		orcPunch = oPower + randInt(1, 5)
		hHealth = hHealth - orcPunch
		x1 := fmt.Sprintf("Орк оказался быстрее тебя и нанёс тебе %d урона.\n У тебя осталось %d здоровья.", orcPunch, hHealth)
		time.Sleep(3 * time.Second)
		fmt.Println(x1)
		fightWithOrc(orcName, heroName, oPower, hPower, oHealth, hHealth, hMoney)
	} else if oHealth <= 0 {
		winningOrc(heroName, hHealthStatic, hPower, hMoney)
	}
	fmt.Println("Ты погиб! :(")
}

func winningOrc(name string, hHealth, hPower, hMoney int) {
	var answer string
	time.Sleep(3 * time.Second)
	fmt.Println("Ты победил!")
	fmt.Println("Твоя награда - 10 золотых монет!")
	fmt.Println("Закончим на сегодня или продолжим?")
	fmt.Println("Варианты: Закончим или Продолжим")
	hMoney += 5
	fmt.Scan(&answer)
	switch {
	case answer == "Продолжим":
		fightWithMobNew(name, hPower, hHealth, hMoney)
	default:
		panic("LEAVE")
	}

}

func goToShop(name string, hHealth, money, power int) {
	var (
		answer string
	)
	fmt.Println("Добро пожаловать в лавку великого алхимика Тимофиуса!\nУ нас не самый большой выбор, мы только открылись\nНо у нас есть всеми любимые булочки путешествеников\nОни помогают восстановить здоровье путникам\nИх цена 5 монет")
	fmt.Println("Будешь покупать?\nВарианты ответа:\nДа \nНет")
	fmt.Printf("Твой баланс:\n%d", money)
	time.Sleep(3 * time.Second)
	if money >= 50 {
		fmt.Println("О, так ты же великий победитель жоп! У меня в жопном кармашке есть специально средтсво для увелечения твоей силы!")
		fmt.Println("Его стоимость 50 монет, хочешь прикупить?")
		fmt.Scan(&answer)
		if answer == "Да" {
			time.Sleep(3 * time.Second)
			money -= 50
			power += 10
			goToShop(name, hHealth, money, power)
		}
	}
	fmt.Scan(&answer)
	if answer == "Да" && money >= 10 {
		money -= 5
		hHealth += 50
	} else if answer == "Да" && money < 5 {
		fmt.Println("Возвращайся, когда наубиваешь орков, нищеброд")
		fightWithMobNew(name, power, hHealth, money)
	}
	fightWithMobNew(name, power, hHealth, money)
}
