/**
 * @fileoverview gRPC-Web generated client stub for unicampus.api.education.v1alpha1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!



const grpc = {};
grpc.web = require('grpc-web');


var google_api_annotations_pb = require('../../../google/api/annotations_pb.js')
const proto = {};
proto.unicampus = {};
proto.unicampus.api = {};
proto.unicampus.api.education = {};
proto.unicampus.api.education.v1alpha1 = require('./generated_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

  /**
   * @private @const {?Object} The credentials to be used to connect
   *    to the server
   */
  this.credentials_ = credentials;

  /**
   * @private @const {?Object} Options for the client
   */
  this.options_ = options;
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!proto.unicampus.api.education.v1alpha1.AdmissionServiceClient} The delegate callback based client
   */
  this.delegateClient_ = new proto.unicampus.api.education.v1alpha1.AdmissionServiceClient(
      hostname, credentials, options);

};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.unicampus.api.education.v1alpha1.Query,
 *   !proto.unicampus.api.education.v1alpha1.School>}
 */
const methodInfo_AdmissionService_ListSchoolsByQuery = new grpc.web.AbstractClientBase.MethodInfo(
  proto.unicampus.api.education.v1alpha1.School,
  /** @param {!proto.unicampus.api.education.v1alpha1.Query} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.unicampus.api.education.v1alpha1.School.deserializeBinary
);


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Query} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.School>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient.prototype.listSchoolsByQuery =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/ListSchoolsByQuery',
      request,
      metadata,
      methodInfo_AdmissionService_ListSchoolsByQuery);
};


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Query} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.School>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient.prototype.listSchoolsByQuery =
    function(request, metadata) {
  return this.delegateClient_.client_.serverStreaming(this.delegateClient_.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/ListSchoolsByQuery',
      request,
      metadata,
      methodInfo_AdmissionService_ListSchoolsByQuery);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.unicampus.api.education.v1alpha1.Critera,
 *   !proto.unicampus.api.education.v1alpha1.School>}
 */
const methodInfo_AdmissionService_ListSchoolsByCritera = new grpc.web.AbstractClientBase.MethodInfo(
  proto.unicampus.api.education.v1alpha1.School,
  /** @param {!proto.unicampus.api.education.v1alpha1.Critera} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.unicampus.api.education.v1alpha1.School.deserializeBinary
);


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Critera} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.School>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient.prototype.listSchoolsByCritera =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/ListSchoolsByCritera',
      request,
      metadata,
      methodInfo_AdmissionService_ListSchoolsByCritera);
};


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Critera} request The request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.School>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient.prototype.listSchoolsByCritera =
    function(request, metadata) {
  return this.delegateClient_.client_.serverStreaming(this.delegateClient_.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/ListSchoolsByCritera',
      request,
      metadata,
      methodInfo_AdmissionService_ListSchoolsByCritera);
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.unicampus.api.education.v1alpha1.School,
 *   !proto.unicampus.api.education.v1alpha1.School>}
 */
const methodInfo_AdmissionService_RegisterSchool = new grpc.web.AbstractClientBase.MethodInfo(
  proto.unicampus.api.education.v1alpha1.School,
  /** @param {!proto.unicampus.api.education.v1alpha1.School} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.unicampus.api.education.v1alpha1.School.deserializeBinary
);


/**
 * @param {!proto.unicampus.api.education.v1alpha1.School} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.unicampus.api.education.v1alpha1.School)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.School>|undefined}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient.prototype.registerSchool =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/RegisterSchool',
      request,
      metadata,
      methodInfo_AdmissionService_RegisterSchool,
      callback);
};


/**
 * @param {!proto.unicampus.api.education.v1alpha1.School} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.unicampus.api.education.v1alpha1.School>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient.prototype.registerSchool =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.registerSchool(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.unicampus.api.education.v1alpha1.School,
 *   !proto.unicampus.api.education.v1alpha1.School>}
 */
const methodInfo_AdmissionService_UpdateSchool = new grpc.web.AbstractClientBase.MethodInfo(
  proto.unicampus.api.education.v1alpha1.School,
  /** @param {!proto.unicampus.api.education.v1alpha1.School} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.unicampus.api.education.v1alpha1.School.deserializeBinary
);


/**
 * @param {!proto.unicampus.api.education.v1alpha1.School} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.unicampus.api.education.v1alpha1.School)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.School>|undefined}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient.prototype.updateSchool =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/UpdateSchool',
      request,
      metadata,
      methodInfo_AdmissionService_UpdateSchool,
      callback);
};


/**
 * @param {!proto.unicampus.api.education.v1alpha1.School} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.unicampus.api.education.v1alpha1.School>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient.prototype.updateSchool =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.updateSchool(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.unicampus.api.education.v1alpha1.School,
 *   !proto.unicampus.api.education.v1alpha1.School>}
 */
const methodInfo_AdmissionService_UnregisterSchool = new grpc.web.AbstractClientBase.MethodInfo(
  proto.unicampus.api.education.v1alpha1.School,
  /** @param {!proto.unicampus.api.education.v1alpha1.School} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.unicampus.api.education.v1alpha1.School.deserializeBinary
);


/**
 * @param {!proto.unicampus.api.education.v1alpha1.School} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.unicampus.api.education.v1alpha1.School)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.School>|undefined}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient.prototype.unregisterSchool =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/UnregisterSchool',
      request,
      metadata,
      methodInfo_AdmissionService_UnregisterSchool,
      callback);
};


/**
 * @param {!proto.unicampus.api.education.v1alpha1.School} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.unicampus.api.education.v1alpha1.School>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient.prototype.unregisterSchool =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.unregisterSchool(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.unicampus.api.education.v1alpha1.Student,
 *   !proto.unicampus.api.education.v1alpha1.Student>}
 */
const methodInfo_AdmissionService_RegisterStudent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.unicampus.api.education.v1alpha1.Student,
  /** @param {!proto.unicampus.api.education.v1alpha1.Student} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.unicampus.api.education.v1alpha1.Student.deserializeBinary
);


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Student} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.unicampus.api.education.v1alpha1.Student)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.Student>|undefined}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient.prototype.registerStudent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/RegisterStudent',
      request,
      metadata,
      methodInfo_AdmissionService_RegisterStudent,
      callback);
};


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Student} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.unicampus.api.education.v1alpha1.Student>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient.prototype.registerStudent =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.registerStudent(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.unicampus.api.education.v1alpha1.Student,
 *   !proto.unicampus.api.education.v1alpha1.Student>}
 */
const methodInfo_AdmissionService_UpdateStudent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.unicampus.api.education.v1alpha1.Student,
  /** @param {!proto.unicampus.api.education.v1alpha1.Student} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.unicampus.api.education.v1alpha1.Student.deserializeBinary
);


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Student} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.unicampus.api.education.v1alpha1.Student)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.Student>|undefined}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient.prototype.updateStudent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/UpdateStudent',
      request,
      metadata,
      methodInfo_AdmissionService_UpdateStudent,
      callback);
};


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Student} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.unicampus.api.education.v1alpha1.Student>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient.prototype.updateStudent =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.updateStudent(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.unicampus.api.education.v1alpha1.Student,
 *   !proto.unicampus.api.education.v1alpha1.Student>}
 */
const methodInfo_AdmissionService_UnregisterStudent = new grpc.web.AbstractClientBase.MethodInfo(
  proto.unicampus.api.education.v1alpha1.Student,
  /** @param {!proto.unicampus.api.education.v1alpha1.Student} request */
  function(request) {
    return request.serializeBinary();
  },
  proto.unicampus.api.education.v1alpha1.Student.deserializeBinary
);


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Student} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.unicampus.api.education.v1alpha1.Student)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.unicampus.api.education.v1alpha1.Student>|undefined}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServiceClient.prototype.unregisterStudent =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/unicampus.api.education.v1alpha1.AdmissionService/UnregisterStudent',
      request,
      metadata,
      methodInfo_AdmissionService_UnregisterStudent,
      callback);
};


/**
 * @param {!proto.unicampus.api.education.v1alpha1.Student} request The
 *     request proto
 * @param {!Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.unicampus.api.education.v1alpha1.Student>}
 *     The XHR Node Readable Stream
 */
proto.unicampus.api.education.v1alpha1.AdmissionServicePromiseClient.prototype.unregisterStudent =
    function(request, metadata) {
  return new Promise((resolve, reject) => {
    this.delegateClient_.unregisterStudent(
      request, metadata, (error, response) => {
        error ? reject(error) : resolve(response);
      });
  });
};


module.exports = proto.unicampus.api.education.v1alpha1;

