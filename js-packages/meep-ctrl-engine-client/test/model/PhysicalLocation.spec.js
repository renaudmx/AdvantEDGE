/**
 * MEEP Controller REST API
 * Copyright (c) 2019 InterDigital Communications, Inc. All rights reserved. The information provided herein is the proprietary and confidential information of InterDigital Communications, Inc. 
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
    // AMD.
    define(['expect.js', '../../src/index'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    factory(require('expect.js'), require('../../src/index'));
  } else {
    // Browser globals (root is window)
    factory(root.expect, root.MeepControllerRestApi);
  }
}(this, function(expect, MeepControllerRestApi) {
  'use strict';

  var instance;

  beforeEach(function() {
    instance = new MeepControllerRestApi.PhysicalLocation();
  });

  var getProperty = function(object, getter, property) {
    // Use getter method if present; otherwise, get the property directly.
    if (typeof object[getter] === 'function')
      return object[getter]();
    else
      return object[property];
  }

  var setProperty = function(object, setter, property, value) {
    // Use setter method if present; otherwise, set the property directly.
    if (typeof object[setter] === 'function')
      object[setter](value);
    else
      object[property] = value;
  }

  describe('PhysicalLocation', function() {
    it('should create an instance of PhysicalLocation', function() {
      // uncomment below and update the code to test PhysicalLocation
      //var instane = new MeepControllerRestApi.PhysicalLocation();
      //expect(instance).to.be.a(MeepControllerRestApi.PhysicalLocation);
    });

    it('should have the property id (base name: "id")', function() {
      // uncomment below and update the code to test the property id
      //var instane = new MeepControllerRestApi.PhysicalLocation();
      //expect(instance).to.be();
    });

    it('should have the property name (base name: "name")', function() {
      // uncomment below and update the code to test the property name
      //var instane = new MeepControllerRestApi.PhysicalLocation();
      //expect(instance).to.be();
    });

    it('should have the property type (base name: "type")', function() {
      // uncomment below and update the code to test the property type
      //var instane = new MeepControllerRestApi.PhysicalLocation();
      //expect(instance).to.be();
    });

    it('should have the property isExternal (base name: "isExternal")', function() {
      // uncomment below and update the code to test the property isExternal
      //var instane = new MeepControllerRestApi.PhysicalLocation();
      //expect(instance).to.be();
    });

    it('should have the property networkLocationsInRange (base name: "networkLocationsInRange")', function() {
      // uncomment below and update the code to test the property networkLocationsInRange
      //var instane = new MeepControllerRestApi.PhysicalLocation();
      //expect(instance).to.be();
    });

    it('should have the property processes (base name: "processes")', function() {
      // uncomment below and update the code to test the property processes
      //var instane = new MeepControllerRestApi.PhysicalLocation();
      //expect(instance).to.be();
    });

  });

}));
