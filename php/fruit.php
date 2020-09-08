<?php
class Fruit
{
    public $name;
    public $color;

    function __construct($name, $color)
    {
        $this->name = $name;
        $this->color = $color;
    }

    function __destruct()
    {
        echo "The fruit is {$this->name}, and the color is {$this->color}\n";
    }
}

class Strawberry extends Fruit
{
    public function message()
    {
        echo "I am $this->name, derived from ", get_parent_class($this) , "\n";
    }
}

$apple = new Fruit("Apple", "Red");
$banana = new Fruit("Banana", "Green");
$strawberry = new Strawberry("strawberry", "Red"); 
$strawberry->message();

// function show_fruit($fruit)
// {
//     echo "The fruit is '$fruit->name', and color is '$fruit->color'\n";
// }

// show_fruit($apple);
// show_fruit($banana);