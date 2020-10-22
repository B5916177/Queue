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
    EntDentist,
    EntDentistFromJSON,
    EntDentistFromJSONTyped,
    EntDentistToJSON,
    EntEmployee,
    EntEmployeeFromJSON,
    EntEmployeeFromJSONTyped,
    EntEmployeeToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntQueueEdges
 */
export interface EntQueueEdges {
    /**
     * 
     * @type {EntDentist}
     * @memberof EntQueueEdges
     */
    dentist?: EntDentist;
    /**
     * 
     * @type {EntEmployee}
     * @memberof EntQueueEdges
     */
    employee?: EntEmployee;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntQueueEdges
     */
    patient?: EntPatient;
}

export function EntQueueEdgesFromJSON(json: any): EntQueueEdges {
    return EntQueueEdgesFromJSONTyped(json, false);
}

export function EntQueueEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntQueueEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'dentist': !exists(json, 'Dentist') ? undefined : EntDentistFromJSON(json['Dentist']),
        'employee': !exists(json, 'Employee') ? undefined : EntEmployeeFromJSON(json['Employee']),
        'patient': !exists(json, 'Patient') ? undefined : EntPatientFromJSON(json['Patient']),
    };
}

export function EntQueueEdgesToJSON(value?: EntQueueEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'dentist': EntDentistToJSON(value.dentist),
        'employee': EntEmployeeToJSON(value.employee),
        'patient': EntPatientToJSON(value.patient),
    };
}

