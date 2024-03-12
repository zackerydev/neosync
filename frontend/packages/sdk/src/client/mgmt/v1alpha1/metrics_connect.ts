// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,import_extension=.js"
// @generated from file mgmt/v1alpha1/metrics.proto (package mgmt.v1alpha1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetMetricCountRequest, GetMetricCountResponse, GetRangedMetricsRequest, GetRangedMetricsResponse } from "./metrics_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service mgmt.v1alpha1.MetricsService
 */
export const MetricsService = {
  typeName: "mgmt.v1alpha1.MetricsService",
  methods: {
    /**
     * Retrieve a timed range of records
     *
     * @generated from rpc mgmt.v1alpha1.MetricsService.GetRangedMetrics
     */
    getRangedMetrics: {
      name: "GetRangedMetrics",
      I: GetRangedMetricsRequest,
      O: GetRangedMetricsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * For the given metric and time range, returns the total count found
     *
     * @generated from rpc mgmt.v1alpha1.MetricsService.GetMetricCount
     */
    getMetricCount: {
      name: "GetMetricCount",
      I: GetMetricCountRequest,
      O: GetMetricCountResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

