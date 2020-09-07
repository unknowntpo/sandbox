<?php
class Fruit
{
    public $name;
    public $color;

    /* Set name and color with __construct */
    function __construct($name, $color)
    {
        $this->name = $name;
        $this->color = $color;
    }
}

$apple = new Fruit("Apple", "Red");
$banana = new Fruit("Banana", "Green");

function show_fruit($fruit)
{
    echo "The fruit is '$fruit->name', and color is '$fruit->color'\n";
}

show_fruit($apple);
show_fruit($banana);