#ifndef ATM_H
#define ATM_H

#include <stdio.h>
#include <unistd.h>

// Different state of ATM machine
typedef enum {
    Idle_State,
    Card_Inserted_State,
    Pin_Eentered_State,
    Option_Selected_State,
    Amount_Entered_State,
} eSystemState;

// Different type events
typedef enum {
    Card_Insert_Event,
    Pin_Enter_Event,
    Option_Selection_Event,
    Amount_Enter_Event,
    Amount_Dispatch_Event
} eSystemEvent;

// Prototype of eventhandlers
eSystemState AmountDispatchHandler(void);
eSystemState EnterAmountHandler(void);
eSystemState OptionSelectionHandler(void);
eSystemState EnterPinHandler(void);
eSystemState InsertCardHandler(void);
eSystemEvent ReadEvent();

void getEvent(eSystemEvent e);

#endif /* ATM_H */
