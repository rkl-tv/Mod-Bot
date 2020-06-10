#include <windows.h>
#include <tlhelp32.h>

DWORD findProcessId(const char *processName) {
    PROCESSENTRY32 processEntry;
    ZeroMemory(&processEntry, sizeof(processEntry));
    processEntry.dwSize = sizeof(processEntry);

    HANDLE processTableSnapshot = CreateToolhelp32Snapshot(TH32CS_SNAPPROCESS, 0);
    if (!Process32First(processTableSnapshot, &processEntry)) {
        CloseHandle(processTableSnapshot);
        return 0;
    }

    do {
        if (0 == strncmp((const char *) processEntry.szExeFile, processName, strlen(processName))) {
            CloseHandle(processTableSnapshot);
            return processEntry.th32ProcessID;
        }
    } while (Process32Next(processTableSnapshot, &processEntry));

    CloseHandle(processTableSnapshot);
    return 0;
}
