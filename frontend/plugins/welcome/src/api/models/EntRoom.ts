/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API Playlist Vidoe
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntRoomEdges,
    EntRoomEdgesFromJSON,
    EntRoomEdgesFromJSONTyped,
    EntRoomEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRoom
 */
export interface EntRoom {
    /**
     * Building holds the value of the "building" field.
     * @type {number}
     * @memberof EntRoom
     */
    building?: number;
    /**
     * 
     * @type {EntRoomEdges}
     * @memberof EntRoom
     */
    edges?: EntRoomEdges;
    /**
     * Floor holds the value of the "floor" field.
     * @type {number}
     * @memberof EntRoom
     */
    floor?: number;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntRoom
     */
    id?: number;
    /**
     * Name holds the value of the "name" field.
     * @type {string}
     * @memberof EntRoom
     */
    name?: string;
}

export function EntRoomFromJSON(json: any): EntRoom {
    return EntRoomFromJSONTyped(json, false);
}

export function EntRoomFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRoom {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'building': !exists(json, 'building') ? undefined : json['building'],
        'edges': !exists(json, 'edges') ? undefined : EntRoomEdgesFromJSON(json['edges']),
        'floor': !exists(json, 'floor') ? undefined : json['floor'],
        'id': !exists(json, 'id') ? undefined : json['id'],
        'name': !exists(json, 'name') ? undefined : json['name'],
    };
}

export function EntRoomToJSON(value?: EntRoom | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'building': value.building,
        'edges': EntRoomEdgesToJSON(value.edges),
        'floor': value.floor,
        'id': value.id,
        'name': value.name,
    };
}

