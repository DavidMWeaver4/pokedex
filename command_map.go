package main
import(
	"fmt"
	"errors"
)

func mapCommand(cfg *config, args[]string) error{
	if cfg.Next == ""{
		return nil
	}
	res, err := ListLocationAreas(cfg.Next)
	if err != nil{
		return err
	}
	if res.Next != nil{
		cfg.Next = *res.Next
	}else{
		cfg.Next = ""
	}
	if res.Previous != nil{
		cfg.Prev = *res.Previous
	}else{
		cfg.Prev = ""
	}
	for _, result := range res.Results{
		fmt.Println(result.Name)
	}

	return nil
}

func mapbCommand(cfg *config, args[]string) error{

	if cfg.Prev == ""{
		fmt.Println("you're on the first page")
		return nil
	}
	res, err := ListLocationAreas(cfg.Prev)
	if err != nil{
		return err
	}
	if res.Next != nil{
		cfg.Next = *res.Next
	}else{
		cfg.Next = ""
	}
	if res.Previous != nil{
		cfg.Prev = *res.Previous
	}else{
		cfg.Prev = ""
	}
	for _, result := range res.Results{
		fmt.Println(result.Name)
	}
	return nil
}
