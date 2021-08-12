#include <stdio.h>
#include <unistd.h>
#include "atm.h"

int main(int argc, char *argv[])
{
    eSystemState eNextState = Idle_State;
    eSystemEvent eNewEvent;
    while (1) {
        // Read system Events
        eNewEvent = ReadEvent();
        printf("new event:");
        getEvent(eNewEvent);

        sleep(1);
        switch (eNextState) {
        case Idle_State: {
            if (Card_Insert_Event == eNewEvent) {
                eNextState = InsertCardHandler();
            }
        } break;
        case Card_Inserted_State: {
            if (Pin_Enter_Event == eNewEvent) {
                eNextState = EnterPinHandler();
            }
        } break;
        case Pin_Eentered_State: {
            if (Option_Selection_Event == eNewEvent) {
                eNextState = OptionSelectionHandler();
            }
        } break;
        case Option_Selected_State: {
            if (Amount_Enter_Event == eNewEvent) {
                eNextState = EnterAmountHandler();
            }
        } break;
        case Amount_Entered_State: {
            if (Amount_Dispatch_Event == eNewEvent) {
                eNextState = AmountDispatchHandler();
            }
        } break;
        default:
            printf("not in any state!\n");
            sleep(1);
            break;
        }
    }
    return 0;
}
