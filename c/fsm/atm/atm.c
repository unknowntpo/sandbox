#include <stdio.h>
#include <unistd.h>
#include "atm.h"

// Prototype of eventhandlers
eSystemState AmountDispatchHandler(void)
{
    printf("amount dispatched !\n");
    sleep(1);
    return Idle_State;
}

eSystemState EnterAmountHandler(void)
{
    printf("amount entered!\n");
    sleep(1);

    return Amount_Entered_State;
}

eSystemState OptionSelectionHandler(void)
{
    printf("option selected!\n");
    sleep(1);

    return Option_Selected_State;
}

eSystemState EnterPinHandler(void)
{
    printf("pin entered!\n");
    sleep(1);

    return Pin_Eentered_State;
}

eSystemState InsertCardHandler(void)
{
    printf("card inserted !\n");
    sleep(1);

    return Card_Inserted_State;
}

// TODO: Use init state to enter the fsm.
eSystemEvent cur = -1;

eSystemEvent ReadEvent()
{
    int length = Amount_Dispatch_Event - Card_Insert_Event;
    cur = (cur + 1) % (length + 1);
    return cur;
}

void getEvent(eSystemEvent e)
{
    switch (e) {
    case Card_Insert_Event:
        printf("Card_Insert_Event\n");
        break;
    case Pin_Enter_Event:
        printf("Pin_Enter_Event\n");
        break;

    case Option_Selection_Event:
        printf("Option_Selection_Event\n");
        break;

    case Amount_Enter_Event:
        printf("Amount_Enter_Event\n");
        break;

    case Amount_Dispatch_Event:
        printf("Amount_Dispatch_Event\n");
        break;
    default:
        printf("no such event!\n");
        break;
    }
}
