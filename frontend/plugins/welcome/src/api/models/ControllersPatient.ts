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
/**
 * 
 * @export
 * @interface ControllersPatient
 */
export interface ControllersPatient {
    /**
     * 
     * @type {string}
     * @memberof ControllersPatient
     */
    age?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatient
     */
    bloodtype?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatient
     */
    gender?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatient
     */
    height?: string;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatient
     */
    name?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersPatient
     */
    prefix?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersPatient
     */
    weight?: string;
}

export function ControllersPatientFromJSON(json: any): ControllersPatient {
    return ControllersPatientFromJSONTyped(json, false);
}

export function ControllersPatientFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersPatient {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'age': !exists(json, 'age') ? undefined : json['age'],
        'bloodtype': !exists(json, 'bloodtype') ? undefined : json['bloodtype'],
        'gender': !exists(json, 'gender') ? undefined : json['gender'],
        'height': !exists(json, 'height') ? undefined : json['height'],
        'name': !exists(json, 'name') ? undefined : json['name'],
        'prefix': !exists(json, 'prefix') ? undefined : json['prefix'],
        'weight': !exists(json, 'weight') ? undefined : json['weight'],
    };
}

export function ControllersPatientToJSON(value?: ControllersPatient | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'age': value.age,
        'bloodtype': value.bloodtype,
        'gender': value.gender,
        'height': value.height,
        'name': value.name,
        'prefix': value.prefix,
        'weight': value.weight,
    };
}

