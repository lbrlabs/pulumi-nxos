// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetIpv4StaticRouteNextHop {
    /**
     * Nexthop address.
     */
    address: string;
    /**
     * Description.
     */
    description: string;
    /**
     * Must match first field in the output of `show intf brief` or `unspecified`. Example: `eth1/1` or `vlan100`.
     */
    interfaceId: string;
    /**
     * Object to be tracked.
     */
    object: number;
    /**
     * Route preference.
     */
    preference: number;
    /**
     * Tag value.
     */
    tag: number;
    /**
     * Nexthop VRF.
     */
    vrfName: string;
}

export interface Ipv4StaticRouteNextHop {
    address: string;
    description?: string;
    interfaceId: string;
    object?: number;
    preference?: number;
    tag?: number;
    vrfName: string;
}

export interface RestChildren {
    /**
     * Class name of the child object.
     */
    className: string;
    /**
     * Map of key-value pairs which represents the attributes of the child object.
     */
    content?: {[key: string]: string};
    /**
     * The relative name of the child object.
     */
    rn: string;
}

export namespace config {
    export interface Devices {
        name: string;
        url: string;
    }

}
