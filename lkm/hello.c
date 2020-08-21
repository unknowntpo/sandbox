#include <linux/init.h>
#include <linux/module.h>
#include <linux/kernel.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("Unknowntpo");
MODULE_DESCRIPTION("A simple linux driver for the BBB");
MODULE_VERSION("0.1");              ///< The version of the module

static char *name = "world";        ///< An example LKM argument -- default value is "world"
module_param(name, charp, S_IRUGO); ///< Param desc. charp = char ptr, S_IRUGO can be read/not changed
MODULE_PARM_DESC(name, "The name to display in /var/log/kern.log");  ///< parameter description

static int __init helloBBB_init(void)
{
    printk(KERN_INFO "EBB: Hello %s from the BBB LKM\n", name);
    return 0;
}

static void __exit helloBBB_exit(void)
{
       printk(KERN_INFO "EBB: Goodbye %s from the BBB LKM!\n", name);
}

module_init(helloBBB_init);
module_exit(helloBBB_exit);

