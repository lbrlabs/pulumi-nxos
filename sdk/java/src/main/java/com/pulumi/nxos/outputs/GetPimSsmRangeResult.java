// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPimSsmRangeResult {
    /**
     * @return A device name from the provider configuration.
     * 
     */
    private @Nullable String device;
    /**
     * @return Group list 1.
     * 
     */
    private String groupList1;
    /**
     * @return Group list 2.
     * 
     */
    private String groupList2;
    /**
     * @return Group list 3.
     * 
     */
    private String groupList3;
    /**
     * @return Group list 4.
     * 
     */
    private String groupList4;
    /**
     * @return The distinguished name of the object.
     * 
     */
    private String id;
    /**
     * @return Prefix list name.
     * 
     */
    private String prefixList;
    /**
     * @return Route map name.
     * 
     */
    private String routeMap;
    /**
     * @return Exclude standard SSM range (232.0.0.0/8).
     * 
     */
    private Boolean ssmNone;
    /**
     * @return VRF name.
     * 
     */
    private String vrfName;

    private GetPimSsmRangeResult() {}
    /**
     * @return A device name from the provider configuration.
     * 
     */
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    /**
     * @return Group list 1.
     * 
     */
    public String groupList1() {
        return this.groupList1;
    }
    /**
     * @return Group list 2.
     * 
     */
    public String groupList2() {
        return this.groupList2;
    }
    /**
     * @return Group list 3.
     * 
     */
    public String groupList3() {
        return this.groupList3;
    }
    /**
     * @return Group list 4.
     * 
     */
    public String groupList4() {
        return this.groupList4;
    }
    /**
     * @return The distinguished name of the object.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Prefix list name.
     * 
     */
    public String prefixList() {
        return this.prefixList;
    }
    /**
     * @return Route map name.
     * 
     */
    public String routeMap() {
        return this.routeMap;
    }
    /**
     * @return Exclude standard SSM range (232.0.0.0/8).
     * 
     */
    public Boolean ssmNone() {
        return this.ssmNone;
    }
    /**
     * @return VRF name.
     * 
     */
    public String vrfName() {
        return this.vrfName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPimSsmRangeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String device;
        private String groupList1;
        private String groupList2;
        private String groupList3;
        private String groupList4;
        private String id;
        private String prefixList;
        private String routeMap;
        private Boolean ssmNone;
        private String vrfName;
        public Builder() {}
        public Builder(GetPimSsmRangeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.device = defaults.device;
    	      this.groupList1 = defaults.groupList1;
    	      this.groupList2 = defaults.groupList2;
    	      this.groupList3 = defaults.groupList3;
    	      this.groupList4 = defaults.groupList4;
    	      this.id = defaults.id;
    	      this.prefixList = defaults.prefixList;
    	      this.routeMap = defaults.routeMap;
    	      this.ssmNone = defaults.ssmNone;
    	      this.vrfName = defaults.vrfName;
        }

        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder groupList1(String groupList1) {
            this.groupList1 = Objects.requireNonNull(groupList1);
            return this;
        }
        @CustomType.Setter
        public Builder groupList2(String groupList2) {
            this.groupList2 = Objects.requireNonNull(groupList2);
            return this;
        }
        @CustomType.Setter
        public Builder groupList3(String groupList3) {
            this.groupList3 = Objects.requireNonNull(groupList3);
            return this;
        }
        @CustomType.Setter
        public Builder groupList4(String groupList4) {
            this.groupList4 = Objects.requireNonNull(groupList4);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder prefixList(String prefixList) {
            this.prefixList = Objects.requireNonNull(prefixList);
            return this;
        }
        @CustomType.Setter
        public Builder routeMap(String routeMap) {
            this.routeMap = Objects.requireNonNull(routeMap);
            return this;
        }
        @CustomType.Setter
        public Builder ssmNone(Boolean ssmNone) {
            this.ssmNone = Objects.requireNonNull(ssmNone);
            return this;
        }
        @CustomType.Setter
        public Builder vrfName(String vrfName) {
            this.vrfName = Objects.requireNonNull(vrfName);
            return this;
        }
        public GetPimSsmRangeResult build() {
            final var o = new GetPimSsmRangeResult();
            o.device = device;
            o.groupList1 = groupList1;
            o.groupList2 = groupList2;
            o.groupList3 = groupList3;
            o.groupList4 = groupList4;
            o.id = id;
            o.prefixList = prefixList;
            o.routeMap = routeMap;
            o.ssmNone = ssmNone;
            o.vrfName = vrfName;
            return o;
        }
    }
}
