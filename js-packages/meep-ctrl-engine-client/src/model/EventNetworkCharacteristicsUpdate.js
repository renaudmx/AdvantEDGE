/**
 * MEEP Controller REST API
 * Copyright (c) 2019  InterDigital Communications, Inc Licensed under the Apache License, Version 2.0 (the \"License\"); you may not use this file except in compliance with the License. You may obtain a copy of the License at      http://www.apache.org/licenses/LICENSE-2.0  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an \"AS IS\" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License. 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'));
  } else {
    // Browser globals (root is window)
    if (!root.MeepControllerRestApi) {
      root.MeepControllerRestApi = {};
    }
    root.MeepControllerRestApi.EventNetworkCharacteristicsUpdate = factory(root.MeepControllerRestApi.ApiClient);
  }
}(this, function(ApiClient) {
  'use strict';




  /**
   * The EventNetworkCharacteristicsUpdate model module.
   * @module model/EventNetworkCharacteristicsUpdate
   * @version 1.0.0
   */

  /**
   * Constructs a new <code>EventNetworkCharacteristicsUpdate</code>.
   * Network Characteristics update Event object
   * @alias module:model/EventNetworkCharacteristicsUpdate
   * @class
   */
  var exports = function() {
    var _this = this;







  };

  /**
   * Constructs a <code>EventNetworkCharacteristicsUpdate</code> from a plain JavaScript object, optionally creating a new instance.
   * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
   * @param {Object} data The plain JavaScript object bearing properties of interest.
   * @param {module:model/EventNetworkCharacteristicsUpdate} obj Optional instance to populate.
   * @return {module:model/EventNetworkCharacteristicsUpdate} The populated <code>EventNetworkCharacteristicsUpdate</code> instance.
   */
  exports.constructFromObject = function(data, obj) {
    if (data) {
      obj = obj || new exports();

      if (data.hasOwnProperty('elementName')) {
        obj['elementName'] = ApiClient.convertToType(data['elementName'], 'String');
      }
      if (data.hasOwnProperty('elementType')) {
        obj['elementType'] = ApiClient.convertToType(data['elementType'], 'String');
      }
      if (data.hasOwnProperty('latency')) {
        obj['latency'] = ApiClient.convertToType(data['latency'], 'Number');
      }
      if (data.hasOwnProperty('latencyVariation')) {
        obj['latencyVariation'] = ApiClient.convertToType(data['latencyVariation'], 'Number');
      }
      if (data.hasOwnProperty('throughput')) {
        obj['throughput'] = ApiClient.convertToType(data['throughput'], 'Number');
      }
      if (data.hasOwnProperty('packetLoss')) {
        obj['packetLoss'] = ApiClient.convertToType(data['packetLoss'], 'Number');
      }
    }
    return obj;
  }

  /**
   * Name of the network element to be updated
   * @member {String} elementName
   */
  exports.prototype['elementName'] = undefined;
  /**
   * Type of the network element to be updated
   * @member {module:model/EventNetworkCharacteristicsUpdate.ElementTypeEnum} elementType
   */
  exports.prototype['elementType'] = undefined;
  /**
   * Latency in ms
   * @member {Number} latency
   */
  exports.prototype['latency'] = undefined;
  /**
   * Latency variation in ms
   * @member {Number} latencyVariation
   */
  exports.prototype['latencyVariation'] = undefined;
  /**
   * Throughput limit
   * @member {Number} throughput
   */
  exports.prototype['throughput'] = undefined;
  /**
   * Packet loss percentage
   * @member {Number} packetLoss
   */
  exports.prototype['packetLoss'] = undefined;


  /**
   * Allowed values for the <code>elementType</code> property.
   * @enum {String}
   * @readonly
   */
  exports.ElementTypeEnum = {
    /**
     * value: "OPERATOR"
     * @const
     */
    "OPERATOR": "OPERATOR",
    /**
     * value: "POA"
     * @const
     */
    "POA": "POA",
    /**
     * value: "SCENARIO"
     * @const
     */
    "SCENARIO": "SCENARIO",
    /**
     * value: "ZONE-INTER-EDGE"
     * @const
     */
    "ZONE-INTER-EDGE": "ZONE-INTER-EDGE",
    /**
     * value: "ZONE-INTER-FOG"
     * @const
     */
    "ZONE-INTER-FOG": "ZONE-INTER-FOG",
    /**
     * value: "ZONE-EDGE-FOG"
     * @const
     */
    "ZONE-EDGE-FOG": "ZONE-EDGE-FOG"  };


  return exports;
}));


