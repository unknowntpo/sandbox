<?php
class Fruit
{
    public $name;
    public $color;

    /* Deal with name */
    public function set_name($name)
    {
        $this->name = $name;
    }
    /* Deal with color */
    public function set_color($color)
    {
        $this->color = $color;
    }

    /* Play with pseudo variables $this */
    public function check_this()
    {
        // print isset($this);              //true,   $this exists
        // print gettype($this);            //Object, $this is an object
        // print is_array($this);           //false,  $this isn't an array
        // $array = get_object_vars($this);    //true,   $this's variables are an array
        // foreach($array as $item) {
        //     //print "$item\n";
        //     printf("%s%s\n", $item, (!next($array)) ? "" : "");
        // }
        // print is_object($this);          //true,   $this is still an object
        // print get_class($this);          //YourProject\YourFile\YourClass
        // print get_parent_class($this);   //YourBundle\YourStuff\YourParentClass
        // print gettype($this->name); //object
        print_r($this);                  //delicious data dump of $this
        // print $this->name;        //access $this variable with ->
        echo "\n";
        return;
    }
}

$apple = new Fruit();
$banana = new Fruit();

/* Set property of apple */
$apple->set_name('Apple');
$apple->set_color('red');

/* Set property of banana */
$banana->set_name('Banana');
$banana->set_color('yellow');

function show_fruit($fruit)
{
    echo "The fruit is '$fruit->name', and color is '$fruit->color'\n";
}

show_fruit($apple);
show_fruit($banana);

$banana->check_this();
