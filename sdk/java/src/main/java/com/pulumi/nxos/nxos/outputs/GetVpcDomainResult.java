// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.nxos.nxos.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetVpcDomainResult {
    private String adminState;
    private String autoRecovery;
    private Integer autoRecoveryInterval;
    private Integer delayRestoreOrphanPort;
    private Integer delayRestoreSvi;
    private Integer delayRestoreVpc;
    private @Nullable String device;
    private Integer domainId;
    private Integer dscp;
    private String fastConvergence;
    private String gracefulConsistencyCheck;
    private String id;
    private String l3PeerRouter;
    private String l3PeerRouterSyslog;
    private Integer l3PeerRouterSyslogInterval;
    private String peerGateway;
    private String peerIp;
    private String peerSwitch;
    private Integer rolePriority;
    private String sysMac;
    private Integer systemPriority;
    private Integer track;
    private String virtualIp;

    private GetVpcDomainResult() {}
    public String adminState() {
        return this.adminState;
    }
    public String autoRecovery() {
        return this.autoRecovery;
    }
    public Integer autoRecoveryInterval() {
        return this.autoRecoveryInterval;
    }
    public Integer delayRestoreOrphanPort() {
        return this.delayRestoreOrphanPort;
    }
    public Integer delayRestoreSvi() {
        return this.delayRestoreSvi;
    }
    public Integer delayRestoreVpc() {
        return this.delayRestoreVpc;
    }
    public Optional<String> device() {
        return Optional.ofNullable(this.device);
    }
    public Integer domainId() {
        return this.domainId;
    }
    public Integer dscp() {
        return this.dscp;
    }
    public String fastConvergence() {
        return this.fastConvergence;
    }
    public String gracefulConsistencyCheck() {
        return this.gracefulConsistencyCheck;
    }
    public String id() {
        return this.id;
    }
    public String l3PeerRouter() {
        return this.l3PeerRouter;
    }
    public String l3PeerRouterSyslog() {
        return this.l3PeerRouterSyslog;
    }
    public Integer l3PeerRouterSyslogInterval() {
        return this.l3PeerRouterSyslogInterval;
    }
    public String peerGateway() {
        return this.peerGateway;
    }
    public String peerIp() {
        return this.peerIp;
    }
    public String peerSwitch() {
        return this.peerSwitch;
    }
    public Integer rolePriority() {
        return this.rolePriority;
    }
    public String sysMac() {
        return this.sysMac;
    }
    public Integer systemPriority() {
        return this.systemPriority;
    }
    public Integer track() {
        return this.track;
    }
    public String virtualIp() {
        return this.virtualIp;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVpcDomainResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String adminState;
        private String autoRecovery;
        private Integer autoRecoveryInterval;
        private Integer delayRestoreOrphanPort;
        private Integer delayRestoreSvi;
        private Integer delayRestoreVpc;
        private @Nullable String device;
        private Integer domainId;
        private Integer dscp;
        private String fastConvergence;
        private String gracefulConsistencyCheck;
        private String id;
        private String l3PeerRouter;
        private String l3PeerRouterSyslog;
        private Integer l3PeerRouterSyslogInterval;
        private String peerGateway;
        private String peerIp;
        private String peerSwitch;
        private Integer rolePriority;
        private String sysMac;
        private Integer systemPriority;
        private Integer track;
        private String virtualIp;
        public Builder() {}
        public Builder(GetVpcDomainResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.adminState = defaults.adminState;
    	      this.autoRecovery = defaults.autoRecovery;
    	      this.autoRecoveryInterval = defaults.autoRecoveryInterval;
    	      this.delayRestoreOrphanPort = defaults.delayRestoreOrphanPort;
    	      this.delayRestoreSvi = defaults.delayRestoreSvi;
    	      this.delayRestoreVpc = defaults.delayRestoreVpc;
    	      this.device = defaults.device;
    	      this.domainId = defaults.domainId;
    	      this.dscp = defaults.dscp;
    	      this.fastConvergence = defaults.fastConvergence;
    	      this.gracefulConsistencyCheck = defaults.gracefulConsistencyCheck;
    	      this.id = defaults.id;
    	      this.l3PeerRouter = defaults.l3PeerRouter;
    	      this.l3PeerRouterSyslog = defaults.l3PeerRouterSyslog;
    	      this.l3PeerRouterSyslogInterval = defaults.l3PeerRouterSyslogInterval;
    	      this.peerGateway = defaults.peerGateway;
    	      this.peerIp = defaults.peerIp;
    	      this.peerSwitch = defaults.peerSwitch;
    	      this.rolePriority = defaults.rolePriority;
    	      this.sysMac = defaults.sysMac;
    	      this.systemPriority = defaults.systemPriority;
    	      this.track = defaults.track;
    	      this.virtualIp = defaults.virtualIp;
        }

        @CustomType.Setter
        public Builder adminState(String adminState) {
            this.adminState = Objects.requireNonNull(adminState);
            return this;
        }
        @CustomType.Setter
        public Builder autoRecovery(String autoRecovery) {
            this.autoRecovery = Objects.requireNonNull(autoRecovery);
            return this;
        }
        @CustomType.Setter
        public Builder autoRecoveryInterval(Integer autoRecoveryInterval) {
            this.autoRecoveryInterval = Objects.requireNonNull(autoRecoveryInterval);
            return this;
        }
        @CustomType.Setter
        public Builder delayRestoreOrphanPort(Integer delayRestoreOrphanPort) {
            this.delayRestoreOrphanPort = Objects.requireNonNull(delayRestoreOrphanPort);
            return this;
        }
        @CustomType.Setter
        public Builder delayRestoreSvi(Integer delayRestoreSvi) {
            this.delayRestoreSvi = Objects.requireNonNull(delayRestoreSvi);
            return this;
        }
        @CustomType.Setter
        public Builder delayRestoreVpc(Integer delayRestoreVpc) {
            this.delayRestoreVpc = Objects.requireNonNull(delayRestoreVpc);
            return this;
        }
        @CustomType.Setter
        public Builder device(@Nullable String device) {
            this.device = device;
            return this;
        }
        @CustomType.Setter
        public Builder domainId(Integer domainId) {
            this.domainId = Objects.requireNonNull(domainId);
            return this;
        }
        @CustomType.Setter
        public Builder dscp(Integer dscp) {
            this.dscp = Objects.requireNonNull(dscp);
            return this;
        }
        @CustomType.Setter
        public Builder fastConvergence(String fastConvergence) {
            this.fastConvergence = Objects.requireNonNull(fastConvergence);
            return this;
        }
        @CustomType.Setter
        public Builder gracefulConsistencyCheck(String gracefulConsistencyCheck) {
            this.gracefulConsistencyCheck = Objects.requireNonNull(gracefulConsistencyCheck);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder l3PeerRouter(String l3PeerRouter) {
            this.l3PeerRouter = Objects.requireNonNull(l3PeerRouter);
            return this;
        }
        @CustomType.Setter
        public Builder l3PeerRouterSyslog(String l3PeerRouterSyslog) {
            this.l3PeerRouterSyslog = Objects.requireNonNull(l3PeerRouterSyslog);
            return this;
        }
        @CustomType.Setter
        public Builder l3PeerRouterSyslogInterval(Integer l3PeerRouterSyslogInterval) {
            this.l3PeerRouterSyslogInterval = Objects.requireNonNull(l3PeerRouterSyslogInterval);
            return this;
        }
        @CustomType.Setter
        public Builder peerGateway(String peerGateway) {
            this.peerGateway = Objects.requireNonNull(peerGateway);
            return this;
        }
        @CustomType.Setter
        public Builder peerIp(String peerIp) {
            this.peerIp = Objects.requireNonNull(peerIp);
            return this;
        }
        @CustomType.Setter
        public Builder peerSwitch(String peerSwitch) {
            this.peerSwitch = Objects.requireNonNull(peerSwitch);
            return this;
        }
        @CustomType.Setter
        public Builder rolePriority(Integer rolePriority) {
            this.rolePriority = Objects.requireNonNull(rolePriority);
            return this;
        }
        @CustomType.Setter
        public Builder sysMac(String sysMac) {
            this.sysMac = Objects.requireNonNull(sysMac);
            return this;
        }
        @CustomType.Setter
        public Builder systemPriority(Integer systemPriority) {
            this.systemPriority = Objects.requireNonNull(systemPriority);
            return this;
        }
        @CustomType.Setter
        public Builder track(Integer track) {
            this.track = Objects.requireNonNull(track);
            return this;
        }
        @CustomType.Setter
        public Builder virtualIp(String virtualIp) {
            this.virtualIp = Objects.requireNonNull(virtualIp);
            return this;
        }
        public GetVpcDomainResult build() {
            final var o = new GetVpcDomainResult();
            o.adminState = adminState;
            o.autoRecovery = autoRecovery;
            o.autoRecoveryInterval = autoRecoveryInterval;
            o.delayRestoreOrphanPort = delayRestoreOrphanPort;
            o.delayRestoreSvi = delayRestoreSvi;
            o.delayRestoreVpc = delayRestoreVpc;
            o.device = device;
            o.domainId = domainId;
            o.dscp = dscp;
            o.fastConvergence = fastConvergence;
            o.gracefulConsistencyCheck = gracefulConsistencyCheck;
            o.id = id;
            o.l3PeerRouter = l3PeerRouter;
            o.l3PeerRouterSyslog = l3PeerRouterSyslog;
            o.l3PeerRouterSyslogInterval = l3PeerRouterSyslogInterval;
            o.peerGateway = peerGateway;
            o.peerIp = peerIp;
            o.peerSwitch = peerSwitch;
            o.rolePriority = rolePriority;
            o.sysMac = sysMac;
            o.systemPriority = systemPriority;
            o.track = track;
            o.virtualIp = virtualIp;
            return o;
        }
    }
}
