/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
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
    EntPatientEdges,
    EntPatientEdgesFromJSON,
    EntPatientEdgesFromJSONTyped,
    EntPatientEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatient
 */
export interface EntPatient {
    /**
     * 
     * @type {EntPatientEdges}
     * @memberof EntPatient
     */
    edges?: EntPatientEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatient
     */
    id?: number;
    /**
     * PatientAge holds the value of the "patient_age" field.
     * @type {number}
     * @memberof EntPatient
     */
    patientAge?: number;
    /**
     * PatientGender holds the value of the "patient_gender" field.
     * @type {string}
     * @memberof EntPatient
     */
    patientGender?: string;
    /**
     * PatientName holds the value of the "patient_name" field.
     * @type {string}
     * @memberof EntPatient
     */
    patientName?: string;
    /**
     * PatientPhone holds the value of the "patient_phone" field.
     * @type {number}
     * @memberof EntPatient
     */
    patientPhone?: number;
}

export function EntPatientFromJSON(json: any): EntPatient {
    return EntPatientFromJSONTyped(json, false);
}

export function EntPatientFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatient {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edges': !exists(json, 'edges') ? undefined : EntPatientEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'patientAge': !exists(json, 'patient_age') ? undefined : json['patient_age'],
        'patientGender': !exists(json, 'patient_gender') ? undefined : json['patient_gender'],
        'patientName': !exists(json, 'patient_name') ? undefined : json['patient_name'],
        'patientPhone': !exists(json, 'patient_phone') ? undefined : json['patient_phone'],
    };
}

export function EntPatientToJSON(value?: EntPatient | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edges': EntPatientEdgesToJSON(value.edges),
        'id': value.id,
        'patient_age': value.patientAge,
        'patient_gender': value.patientGender,
        'patient_name': value.patientName,
        'patient_phone': value.patientPhone,
    };
}

