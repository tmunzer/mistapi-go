import re
import yaml
SPEC_IN = "./mist_openapi/workdir/mist.openapi.yml"
SPEC_OUT = "./mist.sdk.yml"
REG = r"^( *)enum:$"

with open(SPEC_IN, "r") as f_in:
    with open(SPEC_OUT, "w") as f_out:
        f_lines = f_in.readlines()
        for line in f_lines:
            f_out.write(line)
            res = re.match(REG, line)
            if res:
                f_out.write(f"{res.groups()[0]}- \"\"\n")


with open(SPEC_OUT, 'r') as f:
    DATA = yaml.load(f, Loader=yaml.loader.SafeLoader)

del DATA["paths"]["/webhook_example/_alarm_"]
del DATA["components"]["schemas"]["webhook_alarms"]
del DATA["components"]["schemas"]["webhook_alarms_events"]
del DATA["components"]["schemas"]["webhook_alarm_event"]

del DATA["paths"]["/webhook_example/_asset_raw_"]
del DATA["components"]["schemas"]["webhook_asset_raw"]
del DATA["components"]["schemas"]["webhook_asset_raw_events"]
del DATA["components"]["schemas"]["webhook_asset_raw_event"]

del DATA["paths"]["/webhook_example/_audit_"]
del DATA["components"]["schemas"]["webhook_audits"]
del DATA["components"]["schemas"]["webhook_audits_events"]
del DATA["components"]["schemas"]["webhook_audit_event"]

del DATA["paths"]["/webhook_example/_client_info_"]
del DATA["components"]["schemas"]["webhook_client_info"]
del DATA["components"]["schemas"]["webhook_client_info_events"]
del DATA["components"]["schemas"]["webhook_client_info_event"]

del DATA["paths"]["/webhook_example/_client_join_"]
del DATA["components"]["schemas"]["webhook_client_join"]
del DATA["components"]["schemas"]["webhook_client_join_events"]
del DATA["components"]["schemas"]["webhook_client_join_event"]

del DATA["paths"]["/webhook_example/_client_latency_"]
del DATA["components"]["schemas"]["webhook_client_latency"]
del DATA["components"]["schemas"]["webhook_client_latency_events"]
del DATA["components"]["schemas"]["webhook_client_latency_event"]

del DATA["paths"]["/webhook_example/_client_sessions_"]
del DATA["components"]["schemas"]["webhook_client_sessions"]
del DATA["components"]["schemas"]["webhook_client_sessions_events"]
del DATA["components"]["schemas"]["webhook_client_sessions_event"]

del DATA["paths"]["/webhook_example/_device_events_"]
del DATA["components"]["schemas"]["webhook_device_events"]
del DATA["components"]["schemas"]["webhook_device_events_events"]
del DATA["components"]["schemas"]["webhook_device_events_event"]

del DATA["paths"]["/webhook_example/_device_updowns_"]
del DATA["components"]["schemas"]["webhook_device_updowns"]
del DATA["components"]["schemas"]["webhook_device_updowns_events"]
del DATA["components"]["schemas"]["webhook_device_updowns_event"]

del DATA["paths"]["/webhook_example/_discovered_raw_rssi_"]
del DATA["components"]["schemas"]["webhook_discovered_raw_rssi"]
del DATA["components"]["schemas"]["webhook_discovered_raw_rssi_events"]
del DATA["components"]["schemas"]["webhook_discovered_raw_rssi_event"]
del DATA["components"]["schemas"]["webhook_discovered_raw_rssi_event_ap_loc"]

del DATA["paths"]["/webhook_example/_location_"]
del DATA["components"]["schemas"]["webhook_location"]
del DATA["components"]["schemas"]["webhook_location_events"]
del DATA["components"]["schemas"]["webhook_location_event"]
del DATA["components"]["schemas"]["webhook_location_event_wifi_beacon_extended_info"]
del DATA["components"]["schemas"]["webhook_location_event_wifi_beacon_extended_info_items"]

del DATA["paths"]["/webhook_example/_location_asset_"]
del DATA["components"]["schemas"]["webhook_location_asset"]
del DATA["components"]["schemas"]["webhook_location_asset_events"]
del DATA["components"]["schemas"]["webhook_location_asset_event"]

del DATA["paths"]["/webhook_example/_location_centrak_"]
del DATA["components"]["schemas"]["webhook_location_centrak"]
del DATA["components"]["schemas"]["webhook_location_centrak_events"]
del DATA["components"]["schemas"]["webhook_location_centrak_event"]

del DATA["paths"]["/webhook_example/_location_client_"]
del DATA["components"]["schemas"]["webhook_location_client"]
del DATA["components"]["schemas"]["webhook_location_client_events"]
del DATA["components"]["schemas"]["webhook_location_client_event"]

del DATA["paths"]["/webhook_example/_location_sdk_"]
del DATA["components"]["schemas"]["webhook_location_sdk"]
del DATA["components"]["schemas"]["webhook_location_sdk_events"]
del DATA["components"]["schemas"]["webhook_location_sdk_event"]

del DATA["paths"]["/webhook_example/_location_unclient_"]
del DATA["components"]["schemas"]["webhook_location_unclient"]
del DATA["components"]["schemas"]["webhook_location_unclient_events"]
del DATA["components"]["schemas"]["webhook_location_unclient_event"]

del DATA["paths"]["/webhook_example/_nac_accounting_"]
del DATA["components"]["schemas"]["webhook_nac_accounting"]
del DATA["components"]["schemas"]["webhook_nac_accounting_events"]
del DATA["components"]["schemas"]["webhook_nac_accounting_event"]

del DATA["paths"]["/webhook_example/_nac_events_"]
del DATA["components"]["schemas"]["webhook_nac_events"]
del DATA["components"]["schemas"]["webhook_nac_events_events"]
del DATA["components"]["schemas"]["webhook_nac_events_event"]
del DATA["components"]["schemas"]["webhook_nac_events_event_idp_role"]
del DATA["components"]["schemas"]["webhook_nac_events_event_resp_attrs"]

del DATA["paths"]["/webhook_example/_occupancy_alerts_"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts_events"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts_event"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts_event_alert_events"]
del DATA["components"]["schemas"]["webhook_occupancy_alerts_event_alert_events_items"]
del DATA["components"]["schemas"]["webhook_occupancy_alert_type"]

del DATA["paths"]["/webhook_example/_ping_"]
del DATA["components"]["schemas"]["webhook_ping"]
del DATA["components"]["schemas"]["webhook_ping_events"]
del DATA["components"]["schemas"]["webhook_ping_event"]

del DATA["paths"]["/webhook_example/_sdkclient_scan_data"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_events"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_event"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_topic"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_event_scan_data"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_event_scan_data_item"]
del DATA["components"]["schemas"]["webhook_sdkclient_scan_data_event_scan_data_item_band"]

del DATA["paths"]["/webhook_example/_site_sle_"]
del DATA["components"]["schemas"]["webhook_site_sle"]
del DATA["components"]["schemas"]["webhook_site_sle_events"]
del DATA["components"]["schemas"]["webhook_site_sle_event"]
del DATA["components"]["schemas"]["webhook_site_sle_event_sle"]

del DATA["paths"]["/webhook_example/_zone_"]
del DATA["components"]["schemas"]["webhook_zone"]
del DATA["components"]["schemas"]["webhook_zone_events"]
del DATA["components"]["schemas"]["webhook_zone_event"]
del DATA["components"]["schemas"]["webhook_zone_event_trigger"]

with open(SPEC_OUT, 'w') as f:
        yaml.dump({"openapi": DATA["openapi"]}, f)
        yaml.dump({"info": DATA["info"]}, f)
        yaml.dump({"servers": DATA["servers"]}, f)
        yaml.dump({"security": DATA["security"]}, f)
        yaml.dump({"tags": DATA["tags"]}, f)
        yaml.dump({"paths": DATA["paths"]}, f)
        yaml.dump({"components": DATA["components"]}, f)
