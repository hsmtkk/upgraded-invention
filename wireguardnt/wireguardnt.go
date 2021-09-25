package wireguardnt

// https://git.zx2c4.com/wireguard-nt/tree/api/wireguard.h

//go:generate go run github.com/hsmtkk/upgraded-invention/command/mkwinsyscall -output wireguardnt_gen.go wireguardnt.go

/*
typedef struct _WIREGUARD_ADAPTER
{
    HDEVINFO DevInfo;
    SP_DEVINFO_DATA DevInfoData;
    GUID CfgInstanceID;
    WCHAR DevInstanceID[MAX_INSTANCE_ID];
    DWORD LuidIndex;
    DWORD IfType;
    DWORD IfIndex;
    WCHAR Pool[WIREGUARD_MAX_POOL];
    HANDLE LogThread;
    DWORD LogState;
} WIREGUARD_ADAPTER;
*/

type WireGuardAdapter struct{}

//sys CreateAdapter(pool string, name string)(err error) = wireguard.WireGuardCreateAdapter
//sys OpenAdapter(pool string, name string)(handle *WireGuardAdapter, err error) = wireguard.WireGuardOpenAdapter
//sys DeleteAdapter(handle *WireGuardAdapter)(err error) = wireguard.DeleteAdapter
